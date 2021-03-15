/**
 * @Time : 2021/3/15 3:43 下午
 * @Author : MassAdobe
 * @Description: utils
**/
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 3:35 下午
 * @Description: 常量池
**/
const (
	TIME_FORMAT_MS    = "2006-01-02 15:04:05"
	TIME_FORMAT_MONTH = "2006-01-02"
	QUESTION_MARK     = "?"
)

var (
	randCodeSeqCodes = []rune("0123456789")
	randSeqLetters   = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890") // 随机字符串基础值
)

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 6:55 下午
 * @Description: 生成随机时间戳标志位
**/
func RandomTimestampMark() string {
	return fmt.Sprintf("%d%d",
		time.Now().UnixNano(),
		RandInt64(1000, 9999))
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 3:36 下午
 * @Description: 运行当前系统命令
**/
func RunInLinuxWithErr(cmd string) (string, error) {
	result, err := exec.Command(cmd).Output()
	if err != nil {
		fmt.Println(err.Error())
	}
	return strings.TrimSpace(string(result)), err
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:12
 * @Description: md5加密
**/
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 区间随机数；返回int64
**/
func RandInt64(min, max int64) int64 {
	if min >= max || min == 0 || max == 0 {
		return max
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min) + min
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 区间随机数；返回int
**/
func RandInt(min, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min) + min
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 随机字符串
**/
func RandSeq(n int) string {
	b, r := make([]rune, n), rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = randSeqLetters[r.Intn(len(randSeqLetters))]
	}
	return string(b)
}

/**
 * @Author: MassAdobe
 * @TIME: 2021/1/7 11:35 上午
 * @Description: 生成幂等token(保证64位长度)
**/
func RandIdempotentToken(userId int64) string {
	formatInt := strconv.FormatInt(userId, 10)
	return formatInt + RandSeq(64-len(formatInt))
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-05-29 21:47
 * @Description: 生成手机验证码
**/
func RandCodeSeq() string {
	b, r := make([]rune, 6), rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = randCodeSeqCodes[r.Intn(len(randCodeSeqCodes))]
	}
	return string(b)
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 获取当前IP地址
**/
func GetIntranetIp() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 4:33 下午
 * @Description: 获取当前系统IP
**/
func ExternalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("did not connect to network")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}
	return ip
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/21 10:53 上午
 * @Description: 获取请求url上的所有参数
**/
func GetRequestUrlParams(uri string) string {
	if strings.Contains(uri, QUESTION_MARK) {
		return uri[strings.Index(uri, QUESTION_MARK)+1:]
	}
	return ""
}

/**
 * @Author: MassAdobe
 * @TIME: 2020-04-26 21:13
 * @Description: 返回时间字符串
**/
func RtnTmString() (timeStr string) {
	timeStr = time.Now().Format(TIME_FORMAT_MS)
	return
}

/**
 * @Author: MassAdobe
 * @TIME: 2021/3/15 4:48 下午
 * @Description: 根据给定日期返回日期字符串
**/
func FormatDate(time time.Time) (timeStr string) {
	timeStr = time.Format(TIME_FORMAT_MONTH)
	return
}

/**
 * @Author: MassAdobe
 * @TIME: 2020/12/17 4:54 下午
 * @Description: 返回当前时间戳
**/
func RtnCurTime() string {
	return time.Now().Format(TIME_FORMAT_MS)
}
