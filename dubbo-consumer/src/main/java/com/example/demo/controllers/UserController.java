package com.example.demo.controllers;

import com.demo.exp.dto.QueryUserParam;
import com.demo.exp.dto.QueryUserResponse;
import com.demo.exp.service.UserService;
import com.example.demo.dto.Result;
import com.example.demo.dto.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.apache.dubbo.config.annotation.DubboReference;
import org.apache.dubbo.rpc.RpcContext;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/api/user")
@Slf4j
public class UserController {
    @DubboReference(retries = 0, timeout = 1000 * 60, application = "dubbogo-demo", group = "dubbo-service", version = "1.0.0")
    UserService userService;

    @GetMapping("/query")
    @ResponseBody
    public Result QueryUser(){
        QueryUserParam req = new QueryUserParam();
        req.setUserId("666");
        RpcContext rpcContext = RpcContext.getContext();
        rpcContext.setAttachment("app", "dubbo-consumer");
        QueryUserResponse resp = null;
        try{
            resp = userService.QueryUser(req);
        }catch (Throwable e){
            log.error("QueryUser error", e);
            return ResultGenerator.genFailResult(e.getMessage());
        }
        return ResultGenerator.genSuccessResult(resp);
    }
}
