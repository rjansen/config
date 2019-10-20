package testutils

import "time"

func StringPointer(v string) *string { return &v }

func IntPointer(v int) *int { return &v }

func FloatPointer(v float32) *float32 { return &v }

func BoolPointer(v bool) *bool { return &v }

func TimePointer(v time.Time) *time.Time { return &v }

func DurationPointer(v time.Duration) *time.Duration { return &v }
