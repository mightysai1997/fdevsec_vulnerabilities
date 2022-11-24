//import java.security.SecureRandom;
import java.util.Random;
public class Pseudorandom {
public static void main(String[] args){
Random random = new Random(); // Sensitive use of Random
//SecureRandom random = new SecureRandom();
byte bytes[] = new byte[20];
random.nextBytes(bytes);
}
}

