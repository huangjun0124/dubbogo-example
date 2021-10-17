package com.demo.exp.service;

import com.demo.exp.dto.QueryUserParam;
import com.demo.exp.dto.QueryUserResponse;

public interface UserService {
    QueryUserResponse QueryUser(QueryUserParam req);
}
