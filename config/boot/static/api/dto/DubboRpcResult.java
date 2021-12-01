package {{.BasePackage}}.api.dto;

import java.io.Serializable;

public class DubboRpcResult<T> implements Serializable {

    private static final long serialVersionUID = 1L;

    public static final int SUCCESS = 200;

    public static final int FAIL = 400;

    public static final int ERROR = 500;

    private int code;

    private String message;

    private T data;

    public DubboRpcResult() {

    }

    public DubboRpcResult(int code, String message, T data) {
        this.code = code;
        this.message = message;
        this.data = data;
    }

    public static <T> DubboRpcResult<T> custom(int code, String message, T data) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.message = message;
        responseEntity.code = code;
        responseEntity.data = data;
        return responseEntity;
    }

    public static <T> DubboRpcResult<T> custom(int code, String message) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.message = message;
        responseEntity.code = code;
        return responseEntity;
    }

    public static <T> DubboRpcResult<T> customSuccess(String message) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.message = message;
        responseEntity.code = SUCCESS;
        return responseEntity;
    }

    public static <T> DubboRpcResult<T> custom(int code, T data) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.data = data;
        responseEntity.code = code;
        return responseEntity;
    }

    public static <T> DubboRpcResult<T> success(T data) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.data = data;
        responseEntity.code = SUCCESS;
        return responseEntity;
    }

    public static <T> DubboRpcResult<T> fail(String message) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.code = FAIL;
        responseEntity.message = message;
        return responseEntity;
    }

    public static <T> DubboRpcResult<T> failData(String message, T data) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.code = FAIL;
        responseEntity.message = message;
        responseEntity.data = data;
        return responseEntity;
    }

    public static <T> DubboRpcResult<T> error(String message) {
        DubboRpcResult<T> responseEntity = new DubboRpcResult<>();
        responseEntity.code = ERROR;
        responseEntity.message = message;
        return responseEntity;
    }

    public int getCode() {
        return code;
    }

    public String getMessage() {
        return message;
    }

    public T getData() {
        return data;
    }

}
