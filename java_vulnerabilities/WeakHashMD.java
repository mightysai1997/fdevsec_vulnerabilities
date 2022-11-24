import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

/**
 *
 * @author breakthesec
 */
public class WeakHashMD {
    public static String hashMe(String str)
    {
         StringBuffer sb=null;
        try
        {
            MessageDigest md = MessageDigest.getInstance("MD5");  // Sensitive:  SHA is not a standard name, for most security providers it's an alias of SHA-1
           // MessageDigest md1 = MessageDigest.getInstance("SHA-512");  // Sensitive
            md.update(str.getBytes());
            byte byteData[] = md.digest();
            sb= new StringBuffer();
            for (int i = 0; i < byteData.length; i++)
            {
             sb.append(Integer.toString((byteData[i] & 0xff) + 0x100, 16).substring(1));
            }
        }
        catch(NoSuchAlgorithmException e)
        {

        }
        return sb.toString();
    }
}

