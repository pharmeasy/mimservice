package com.diagnostics.mimservice.controller;

import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.CrossOrigin;
import org.springframework.web.bind.annotation.RestController;
@RestController
@CrossOrigin
public class HealthController {
    @GetMapping("/status")
    public ResponseEntity<Object> get() throws Exception {
        return ResponseEntity.ok().body("OK");
    }
}
