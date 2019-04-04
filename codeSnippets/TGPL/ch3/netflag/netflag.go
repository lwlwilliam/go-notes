package netflag

type Flags uint

const (
	FlagUp Flags = 1 << iota	// 1 << 0: 0000 0001
	FlagBroadcast				// 1 << 1: 0000 0010
	FlagLoopback				// 1 << 2: 0000 0100
	FlagPointToPoint			// 1 << 3: 0000 1000
	FlagMulticast				// 1 << 4: 0001 0000
)

// 判断是否为 FlagUp 标志
func IsUp(v Flags) bool  {
	return v&FlagUp == FlagUp
}

// 把 FlagUp 位关闭
func TurnDown(v *Flags)  {
	*v &^= FlagUp
}

// 设置 FlagBroadcast 位
func SetBroadcast(v *Flags)  {
	*v |= FlagBroadcast
}

// 判断是否有 cast 标志
func IsCast(v Flags) bool {
	return v&(FlagBroadcast|FlagMulticast) != 0
}
