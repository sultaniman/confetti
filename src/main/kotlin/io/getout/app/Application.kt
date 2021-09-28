package io.getout.app

import org.springframework.boot.autoconfigure.SpringBootApplication
import org.springframework.boot.runApplication
import org.springframework.scheduling.annotation.EnableAsync
import org.springframework.scheduling.annotation.EnableScheduling

@EnableAsync
@EnableScheduling
@SpringBootApplication(scanBasePackageClasses = [Application::class])
class Application

fun main(args: Array<String>) {
	runApplication<Application>(*args)
}
