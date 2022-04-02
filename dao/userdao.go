package dao

import (
	"goitems/model"
	"goitems/utils"
)

//验证用户

func CheckUserNameAndPassword(name string, age int) (*model.User, error) {
	//写sql语句
	sqlStr := "select id, name, age, gender from test where name=? and age=?"
	//执行
	row := utils.Db.QueryRow(sqlStr, name, age)
	user := &model.User{}

	row.Scan(&user.Id, &user.Name, &user.Age, &user.Gender) //注意大写
	return user, nil
}

// //验证用户名字,是否存在

func CheckUserName(name string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id, name, age, gender from test where name = ? "
	//执行
	row := utils.Db.QueryRow(sqlStr, name)
	user := &model.User{}
	row.Scan(&user.Id, &user.Name, &user.Age, &user.Gender)
	//报错以后会返回其他的东西,所以暂时先不动
	// errNotFound := row.Scan(&user.Id, &user.Name, &user.Age, &user.Gender)
	// if errNotFound != nil {
	// 	fmt.Println("errNotFound---", errNotFound)
	// 	return nil, errNotFound
	// }
	return user, nil
}

//插入数据

func SaveUser(name string, age int, gender string) error {
	//写sql语句
	sqlStr := "insert into test(name,age,gender) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, name, age, gender)
	if err != nil {
		return err
	}
	return nil
}

// gender
// 英 [ˈdʒendə(r)]   美 [ˈdʒendər]
// n.
// 性别;性(阳性、阴性和中性，不同的性有不同的词尾等)
// 复数： genders
