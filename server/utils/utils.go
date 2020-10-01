package utils

import (
	"cloudChat/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//将这些方法关联到Transfer（传输者）结构体中
type Transfer struct {
	Conn net.Conn //链接
	Buf [8064]byte //这是传输时的缓冲
}


//将读取数据包
func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	//buf := make([]byte, 1024*4)
	//conn.Read 在conn 没有被关闭时 才会被堵塞，关闭则不会堵塞
	_, err = this.Conn.Read(this.Buf[:4]) //没有读到东西会堵塞
	if err != nil {
		return
	}
	//fmt.Println("读到的buf=", buf[:4])

	//根据buf[:4] 转成一个uint32类型
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen 读取消息内容(将pkgLen长度的内容读到buf里)
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if uint32(n) != pkgLen || err != nil {
		err = errors.New("read pkg body err")
		return
	}

	//把PkgLen 反序列化成 -> message.Message类型（mes）
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarshal(buf[:pkgLen], &mes) err = ", err)
		return
	}

	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[:4], pkgLen)
	//发送长度 n 为字节长度
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err=", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if uint32(n) != pkgLen || err != nil {
		fmt.Println("conn.Write(buf) err=", err)
		return
	}
	return
}