package record

/**
 * @author wangleij52
 * @date 2022/04/21/18:47
 */

type Protocol interface {
	Hit(data [][]byte) (name []byte, ok bool)
}
