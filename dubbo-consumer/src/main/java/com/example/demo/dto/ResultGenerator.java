package com.example.demo.dto;

import org.springframework.util.StringUtils;

public class ResultGenerator {
    private static final String DEFAULT_SUCCESS_MESSAGE = "success";
    private static final String DEFAULT_FAIL_MESSAGE = "fail";
    private static final int RESULT_CODE_SUCCESS = 200;
    private static final  int RESULT_CODE_SERVER_ERROR = 500;

    public static Result genSuccessResult(){
        Result res = new Result();
        res.setResultCode(RESULT_CODE_SUCCESS);
        res.setMessage(DEFAULT_SUCCESS_MESSAGE);
        return res;
    }

    public static Result genSuccessResult(String msg){
        Result res = new Result();
        res.setResultCode(RESULT_CODE_SUCCESS);
        res.setMessage(msg);
        return res;
    }

    public static Result genSuccessResult(Object data){
        Result res = new Result();
        res.setResultCode(RESULT_CODE_SUCCESS);
        res.setMessage(DEFAULT_SUCCESS_MESSAGE);
        res.setData(data);
        return res;
    }

    public static Result genFailResult(String msg){
        Result res = new Result();
        res.setResultCode(RESULT_CODE_SERVER_ERROR);
        if(!StringUtils.hasLength(msg)){
            res.setMessage(DEFAULT_FAIL_MESSAGE);
        }else{
            res.setMessage(msg);
        }
        return res;
    }

    public static Result genErrorResult(int code, String msg){
        Result res = new Result();
        res.setResultCode(code);
        res.setMessage(msg);
        return res;
    }
}
