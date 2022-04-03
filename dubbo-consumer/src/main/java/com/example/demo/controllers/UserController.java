package com.example.demo.controllers;

import com.demo.exp.dto.QueryUserParam;
import com.demo.exp.dto.QueryUserResponse;
import com.demo.exp.service.UserService;
import com.example.demo.dto.Result;
import com.example.demo.dto.ResultGenerator;
import lombok.extern.slf4j.Slf4j;
import org.apache.dubbo.config.annotation.DubboReference;
import org.apache.dubbo.rpc.RpcContext;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api/user")
@Slf4j
public class UserController {
    @DubboReference(retries = 0, timeout = 1000 * 60)
    UserService userService;

    @GetMapping("/query")
    @ResponseBody
    public Result QueryUser(@RequestParam String userId){
        QueryUserParam req = new QueryUserParam();
        req.setUserId(userId);
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
