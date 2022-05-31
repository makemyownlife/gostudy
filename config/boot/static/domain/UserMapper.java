package {{.BasePackage}}.domain.mapper;

import {{.BasePackage}}.domain.po.User;
import org.springframework.stereotype.Repository;

@Repository
public interface UserMapper {

    User getUserById(Long id);

}
