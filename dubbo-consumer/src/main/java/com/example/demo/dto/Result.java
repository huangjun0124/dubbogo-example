package com.example.demo.dto;

import com.fasterxml.jackson.annotation.JsonInclude;
import lombok.Getter;
import lombok.Setter;

import java.io.Serializable;

@Getter
@Setter
@JsonInclude(JsonInclude.Include.NON_NULL)
public class Result<T> implements Serializable {
    private static final long serialVersionUID = 1L;
    // 业务码code
    private int resultCode;
    // 返回信息
    private String message;
    // 返回的数据
    private T data;

    public Result(){}

    public Result(int code, String message){
        this.resultCode = code;
        this.message = message;
    }

    @Override
    public String toString(){
        return "Result{" +
                "resultCode=" + resultCode +
                ", message='" + message + "'" +
                ", data=" + data +
                "}";
    }

}
