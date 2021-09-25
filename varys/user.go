package varys

type UserInfo struct {
	Name string
}

func GetInfoByUID(int64)(*UserInfo, error){
	return &UserInfo{}, nil
}
