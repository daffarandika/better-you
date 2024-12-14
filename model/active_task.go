package model

type ActiveTask struct {
	ID		int
	TaskID	int	
	Task	Task	`gorm:"foreignKey:TaskID;references:ID"`
	Done	bool
}
