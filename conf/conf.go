package conf

//读取conf.ini，并将数据存放到全局变量中，供其他模块使用
import (
	"fmt"
	"strings"
	"todo_list/model"

	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	Port       string
	DbUser     string
	DbPassWord string
	DbName     string
)

func Init() {
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("load init fail", err)
	}

	LoadServer(file)
	LoadMysql(file)

	path := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", Port, ")/", DbName, "?charset=utf8mb4&parseTime=True"}, "")
	//调用model模块中的函数，进行数据库初始化
	model.DataBase(path)

}

// 配置信息赋值给全局变量
func LoadServer(file *ini.File) {
	AppMode = file.Section("service").Key("AppMode").String()
	HttpPort = file.Section("service").Key("HttpPort").String()
	fmt.Println(AppMode, HttpPort)
}
func LoadMysql(file *ini.File) {
	//Db都为全局变量，读取config.ini的配置信息，并存入全局变量中
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	Port = file.Section("mysql").Key("Port").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()

}
