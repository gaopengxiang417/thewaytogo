public class linuxtest{
	public static void main(String[] args){
	Runtime runtime = Runtime.getRuntime();
    String[] strings = {"/bin/sh","-c","ls -al"};
    Process process = runtime.exec(strings);

    process.waitFor();

    BufferedReader bufferedReader = new BufferedReader(new InputStreamReader(process.getInputStream()));
    String s = null;
    while ((s = bufferedReader.readLine()) != null) {
        System.out.println(s);
    }
}
}


