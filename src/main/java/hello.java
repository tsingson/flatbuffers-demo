import Serializer.Event;
import java.io.File;
import java.io.RandomAccessFile;
import java.nio.ByteBuffer;
public class hello {
    public static void main(String[] args) {
        byte[] data = null;
        File file = new File("/Users/qinshen/go/src/github.com/tsingson/flatbuffers-demo/test.bin");
        RandomAccessFile f = null;
        try {
            f = new RandomAccessFile(file, "r");
            data = new byte[(int) f.length()];
            f.readFully(data);
            f.close();
        } catch (java.io.IOException e) {
            System.out.println("FlatBuffers test: couldn't read file");
            return;
        }

        ByteBuffer bb = ByteBuffer.wrap(data);

        Event e = Event.getRootAsEvent(bb);
        System.out.println(e.parentWindowId());
        System.out.println(e.windowId());
    }
}