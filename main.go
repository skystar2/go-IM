package main

// Todo 解析配置文件
// Todo 初始化log
func init() {
	// fmt.Println("。。。。。")
}

func main() {
	server := NewServer("127.0.0.1", 8888, 30)
	server.Start()
}
