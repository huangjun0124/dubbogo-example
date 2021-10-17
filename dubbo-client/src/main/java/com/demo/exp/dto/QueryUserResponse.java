package com.demo.exp.dto;

import lombok.Data;

import java.io.Serializable;

@Data
public class QueryUserResponse implements Serializable {
    private String userId;
    private String userName;
    private String birthDate;
    private long age;
    private boolean isDead;
}
