package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//定义一个 UserDao 的结构体(完成 user 结构体的各种方法)
type UserDao struct {
	pool *redis.Pool
}

//全局的 UserDao 实例变量，在需要使用redis操作是，直接使用该实例（可以提高redis的使用效率）
var (
	MyUserDao *UserDao
)

//1. 根据 userID 返回 User 实例+ err (校验用户id存在)
func (this *UserDao) getUserById(conn redis.Conn, userId int) (user *User, err error) {
	//通过 userId 去 redis 查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", userId))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示在 users 哈希中，没有找到对应的userId
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}

	//将 res 反序列化的 User 实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
	}

	return
}

//2. 完成登录的校验
//2.1 如果用户的 id 与 pwd 都正确，则返回一个 user 实例
//2.2 如果用户的 id 与 pwd 有错误， 则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	//先从 UserDao 的连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()

	//获取user实例，既获取user存在redis中的基本信息
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//交接用户的密码是否正确( user.UserPwd 是从redis中取出来的密码，userPwd是用户输入的密码)
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
	}

	return
}

//3. 使用工厂模式，创建一个 UserDao 实例(在服务器连接是接通连接池)
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	//创建一个 UserDao 实例
	userDao = &UserDao{
		pool: pool,
	}

	return
}
