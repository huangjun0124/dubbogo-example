package com.demo.exp.dto;

import lombok.Data;

import java.io.Serializable;

@Data
public class QueryUserParam implements Serializable {
    private String userName;
    private String userId;
}
