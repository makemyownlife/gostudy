package {{.BasePackage}}.service;

import {{.BasePackage}}.domain.mapper.UserMapper;
import {{.BasePackage}}.domain.po.User;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class UserService {

    @Autowired
    private UserMapper userMapper;

    public User getUserById(Long id) {
        return userMapper.getUserById(id);
    }

}
