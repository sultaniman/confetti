package io.getout.app.system.controllers

import org.springframework.http.ResponseEntity
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController

@RestController
@RequestMapping
class Health {
    @GetMapping(value = ["/api/health"])
    fun health() : ResponseEntity<String> {
        return ResponseEntity.ok("ok")
    }
}
