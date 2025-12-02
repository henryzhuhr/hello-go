/**
 * 用户查询服务：实现“布隆过滤器→Redis→数据库”三级过滤
 * 防穿透核心设计
 * 布隆过滤器前置拦截：对“绝对不存在”的ID直接返回，避免请求到达Redis和数据库，从源头减少无效流量。
 * 空值缓存机制：针对布隆过滤器的“假阳性”（判断存在但实际不存在），缓存空值1分钟，防止同一ID重复穿透数据库。
 * 三级过滤顺序：严格遵循“布隆过滤器→Redis→数据库”的顺序，层层过滤高成本操作，最大化减轻数据库压力。
 */
public class UserService {
    // 注入依赖（实际项目中用Spring等框架注入）
    private BloomFilter bloomFilter;
    private RedisClient redisClient;
    private UserDao userDao;

    // 缓存过期时间（30分钟，可根据业务调整）
    private static final int CACHE_EXPIRE_SECONDS = 30 * 60;

    /**
     * 根据用户ID查询用户信息（核心方法）
     */
    public String getUserById(String userId) {
        // 1. 第一级：布隆过滤器预判断（绝对拦截不存在的ID）
        if (!bloomFilter.contains(userId)) {
            System.out.println("布隆过滤器拦截：ID不存在 → " + userId);
            return null; // 直接返回，不进入后续流程
        }

        // 2. 第二级：Redis缓存查询（命中则直接返回，避免查库）
        String userInfo = redisClient.get("user:" + userId);
        if (userInfo != null) {
            System.out.println("Redis命中：返回用户信息 → " + userId);
            return userInfo;
        }

        // 3. 第三级：数据库查询（缓存未命中时查库，并更新缓存）
        userInfo = userDao.findUserById(userId);
        if (userInfo != null) {
            // 数据库存在该用户：缓存真实数据
            redisClient.set("user:" + userId, userInfo, CACHE_EXPIRE_SECONDS);
            System.out.println("数据库命中：缓存并返回用户信息 → " + userId);
        } else {
            // 数据库不存在该用户（布隆过滤器假阳性）：缓存空值（短期）
            redisClient.set("user:" + userId, "", 60); // 空值缓存1分钟，避免重复穿透
            System.out.println("数据库未命中：缓存空值 → " + userId);
        }

        return userInfo;
    }
}
