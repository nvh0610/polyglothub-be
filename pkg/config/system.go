package config

func MaxFlashCardDaily() int {
	return IntEnvF("MAX_FLASHCARD_DAILY", 20)
}

func CronJobFetchFlashCardDaily() string {
	return StringEnvF("CRON_JOB_FETCH_FLASHCARD_DAILY", "@every 0h3m0s")
}
