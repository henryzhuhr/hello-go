package main

import (
	"fmt"
	"math/rand"
	"time"

	bloomfilter "github.com/henryzhuhr/hello-go/internal/algorithm/bloom-filter"
)

// UserService 用户查询服务：实现"布隆过滤器→Redis→数据库"三级过滤
// 防穿透核心设计：
// 1. 布隆过滤器前置拦截：对"绝对不存在"的ID直接返回，避免请求到达Redis和数据库，从源头减少无效流量。
// 2. 空值缓存机制：针对布隆过滤器的"假阳性"（判断存在但实际不存在），缓存空值1分钟，防止同一ID重复穿透数据库。
// 3. 三级过滤顺序：严格遵循"布隆过滤器→Redis→数据库"的顺序，层层过滤高成本操作，最大化减轻数据库压力。
type UserService struct {
	bloomFilter bloomfilter.BloomFilter[string]
	redisClient RedisClient
	userDao     UserDao
}

// RedisClient Redis客户端接口
type RedisClient interface {
	// Get 获取缓存值
	Get(key string) (string, error)

	// Set 设置缓存值（带过期时间）
	Set(key string, value string, expireSeconds int) error
}

// UserDao 用户数据访问接口
type UserDao interface {
	// FindUserByID 根据用户ID查询用户信息
	FindUserByID(userID string) (string, error)
}

const (
	// CacheExpireSeconds 缓存过期时间（30分钟，可根据业务调整）
	CacheExpireSeconds = 30 * 60
	// EmptyCacheExpireSeconds 空值缓存过期时间（1分钟）
	EmptyCacheExpireSeconds = 60
)

// NewUserService 创建用户服务实例
func NewUserService(
	bloomFilter bloomfilter.BloomFilter[string],
	redisClient RedisClient,
	userDao UserDao,
) *UserService {
	return &UserService{
		bloomFilter: bloomFilter,
		redisClient: redisClient,
		userDao:     userDao,
	}
}

// GetUserByID 根据用户ID查询用户信息（核心方法）
func (s *UserService) GetUserByID(userID string) (string, error) {
	// 1. 第一级：布隆过滤器预判断（绝对拦截不存在的ID）
	if !s.bloomFilter.Contains(userID) {
		return "", nil // 直接返回，不进入后续流程
	}

	// 2. 第二级：Redis缓存查询（命中则直接返回，避免查库）
	cacheKey := "user:" + userID
	userInfo, err := s.redisClient.Get(cacheKey)
	if err == nil && userInfo != "" {
		return userInfo, nil
	}

	// 如果缓存了空值，直接返回（避免重复穿透）
	if err == nil && userInfo == "" {
		return "", nil
	}

	// 3. 第三级：数据库查询（缓存未命中时查库，并更新缓存）
	userInfo, err = s.userDao.FindUserByID(userID)
	if err != nil {
		return "", err
	}

	if userInfo != "" {
		// 数据库存在该用户：缓存真实数据
		if err := s.redisClient.Set(cacheKey, userInfo, CacheExpireSeconds); err != nil {
			fmt.Printf("Redis缓存失败: %v\n", err)
		}
	} else {
		// 数据库不存在该用户（布隆过滤器假阳性）：缓存空值（短期）
		if err := s.redisClient.Set(cacheKey, "", EmptyCacheExpireSeconds); err != nil {
			fmt.Printf("Redis缓存空值失败: %v\n", err)
		}
	}

	return userInfo, nil
}

// ===== 以下是模拟实现，用于演示 =====

// mockRedisClient 模拟Redis客户端
type mockRedisClient struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	value      string
	expireTime time.Time
}

// NewMockRedisClient 创建模拟Redis客户端
func NewMockRedisClient() RedisClient {
	return &mockRedisClient{
		cache: make(map[string]cacheEntry),
	}
}

// Get implements [RedisClient.Get]
func (m *mockRedisClient) Get(key string) (string, error) {
	entry, exists := m.cache[key]
	if !exists {
		return "", fmt.Errorf("key not found")
	}

	// 检查是否过期
	if time.Now().After(entry.expireTime) {
		delete(m.cache, key)
		return "", fmt.Errorf("key expired")
	}

	return entry.value, nil
}

// Set implements [RedisClient.Set]
func (m *mockRedisClient) Set(key string, value string, expireSeconds int) error {
	m.cache[key] = cacheEntry{
		value:      value,
		expireTime: time.Now().Add(time.Duration(expireSeconds) * time.Second),
	}
	return nil
}

// mockUserDao 模拟用户数据访问
type mockUserDao struct {
	users      map[string]string
	queryCount int           // 查询次数统计
	totalTime  time.Duration // 总查询时间
}

// NewMockUserDao 创建模拟用户数据访问对象
func NewMockUserDao() *mockUserDao {
	return &mockUserDao{
		users: make(map[string]string),
	}
}

// AddUser 添加用户（用于批量初始化）
func (m *mockUserDao) AddUser(userID string, userInfo string) {
	m.users[userID] = userInfo
}

// FindUserByID implements [UserDao.FindUserByID]
func (m *mockUserDao) FindUserByID(userID string) (string, error) {
	start := time.Now()

	// 模拟数据库查询延迟（1-3毫秒）
	time.Sleep(time.Millisecond * time.Duration(1+rand.Intn(3)))

	userInfo, exists := m.users[userID]

	m.queryCount++
	m.totalTime += time.Since(start)

	if !exists {
		return "", nil
	}
	return userInfo, nil
}

// GetStats 获取统计信息
func (m *mockUserDao) GetStats() (count int, avgTime time.Duration) {
	if m.queryCount == 0 {
		return 0, 0
	}
	return m.queryCount, m.totalTime / time.Duration(m.queryCount)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// ===== 大数据量性能测试 =====
	fmt.Println("===== 用户查询服务性能测试 =====\n")

	// 配置参数
	const (
		totalUsers    = 100000 // 总用户数
		testQueries   = 50000  // 测试查询次数
		invalidRatio  = 0.5    // 无效ID查询比例（50%）
		bloomFilterFP = 0.01   // 布隆过滤器误判率（1%）
	)

	fmt.Printf("测试配置:\n")
	fmt.Printf("  - 总用户数: %d\n", totalUsers)
	fmt.Printf("  - 测试查询次数: %d\n", testQueries)
	fmt.Printf("  - 无效ID查询比例: %.0f%%\n", invalidRatio*100)
	fmt.Printf("  - 布隆过滤器误判率: %.1f%%\n\n", bloomFilterFP*100)

	// 1. 初始化布隆过滤器
	fmt.Println("正在初始化布隆过滤器...")
	bf := bloomfilter.NewBloomFilter(uint64(totalUsers), bloomFilterFP, func(s string) []byte {
		return []byte(s)
	})

	// 2. 初始化数据库和Redis
	userDao := NewMockUserDao()
	redisClient := NewMockRedisClient()

	// 3. 批量生成用户数据并添加到布隆过滤器
	fmt.Printf("正在生成 %d 个用户数据...\n", totalUsers)
	for i := 1; i <= totalUsers; i++ {
		userID := fmt.Sprintf("user_%d", i)
		userInfo := fmt.Sprintf(`{"id":"%s","name":"用户%d","age":%d}`, userID, i, 20+i%50)

		userDao.AddUser(userID, userInfo)
		bf.Add(userID)
	}

	// 4. 创建用户服务
	userService := NewUserService(bf, redisClient, userDao)

	fmt.Println("\n===== 开始性能测试 =====\n")

	// 统计变量
	var (
		bloomFilterBlocked = 0 // 布隆过滤器拦截次数
		redisHit           = 0 // Redis命中次数
		dbQuery            = 0 // 数据库查询次数
		dbHit              = 0 // 数据库命中次数
		dbMiss             = 0 // 数据库未命中次数（假阳性）
	)

	// 5. 执行测试查询
	startTime := time.Now()

	for i := 0; i < testQueries; i++ {
		var userID string

		// 按照比例生成有效/无效ID
		if rand.Float64() < invalidRatio {
			// 生成不存在的ID（范围外）
			userID = fmt.Sprintf("user_%d", totalUsers+1+rand.Intn(totalUsers))
		} else {
			// 生成存在的ID
			userID = fmt.Sprintf("user_%d", 1+rand.Intn(totalUsers))
		}

		// 查询前的数据库查询次数
		prevDBCount, _ := userDao.GetStats()

		// 执行查询
		_, err := userService.GetUserByID(userID)
		if err != nil {
			continue
		}

		// 查询后的数据库查询次数
		currDBCount, _ := userDao.GetStats()

		// 统计结果
		if currDBCount > prevDBCount {
			// 查询了数据库
			dbQuery++

			// 检查缓存中的值来判断是否命中
			cacheValue, err := redisClient.Get("user:" + userID)
			if err == nil {
				if cacheValue != "" {
					dbHit++
				} else {
					dbMiss++ // 假阳性
				}
			}
		} else {
			// 没有查询数据库
			_, err := redisClient.Get("user:" + userID)
			if err == nil {
				// Redis命中
				redisHit++
			} else {
				// 布隆过滤器拦截
				bloomFilterBlocked++
			}
		}
	}

	totalTime := time.Since(startTime)
	dbQueryCount, avgDBTime := userDao.GetStats()

	// 6. 输出性能报告
	fmt.Println("===== 性能测试报告 =====\n")

	fmt.Printf("【总体性能】\n")
	fmt.Printf("  总查询次数: %d\n", testQueries)
	fmt.Printf("  总耗时: %v\n", totalTime)
	fmt.Printf("  平均每次查询: %.2f ms\n", float64(totalTime.Microseconds())/float64(testQueries)/1000)
	fmt.Printf("  QPS: %.0f 次/秒\n\n", float64(testQueries)/totalTime.Seconds())

	fmt.Printf("【三级过滤统计】\n")
	fmt.Printf("  布隆过滤器拦截: %d 次 (%.1f%%)\n",
		bloomFilterBlocked, float64(bloomFilterBlocked)*100/float64(testQueries))
	fmt.Printf("  Redis缓存命中: %d 次 (%.1f%%)\n",
		redisHit, float64(redisHit)*100/float64(testQueries))
	fmt.Printf("  数据库查询: %d 次 (%.1f%%)\n",
		dbQuery, float64(dbQuery)*100/float64(testQueries))
	fmt.Printf("    - 数据库命中: %d 次\n", dbHit)
	fmt.Printf("    - 数据库未命中(假阳性): %d 次\n\n", dbMiss)

	fmt.Printf("【数据库保护效果】\n")
	fmt.Printf("  实际数据库查询次数: %d\n", dbQueryCount)
	fmt.Printf("  数据库查询平均耗时: %v\n", avgDBTime)
	fmt.Printf("  如无保护预计查询次数: %d\n", testQueries)
	fmt.Printf("  减少查询次数: %d (%.1f%%)\n",
		testQueries-dbQueryCount, float64(testQueries-dbQueryCount)*100/float64(testQueries))

	if dbMiss > 0 {
		fmt.Printf("\n【布隆过滤器准确性】\n")
		fmt.Printf("  实际误判率: %.2f%% (%d/%d)\n",
			float64(dbMiss)*100/float64(bloomFilterBlocked+dbQuery), dbMiss, bloomFilterBlocked+dbQuery)
		fmt.Printf("  理论误判率: %.2f%%\n", bloomFilterFP*100)
	}

	fmt.Println("\n===== 测试完成 =====")
}
