package global

import (
	"errors"
	log "github.com/sirupsen/logrus"
)

// DeleteModelById 根据id删除对应的模型
func DeleteModelById(id string) (err error) {
	// 删除modelMap中的内容
	m, ok := ModelIdMap[id]
	if !ok {
		log.Warnln(id, "not exists")
		return errors.New("id not exists")
	}
	m.RemoveDir()
	delete(ModelIdMap, id)
	return nil
}
