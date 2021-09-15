package de.jcup.java_object_behaviour;

public class SingleObjectUsedInMultipleThreadsBlockingTestcase {
    public static void main(String[] args) {
        TheObject obj = new TheObject();
        
        startThread(obj,"t1",500);
        startThread(obj,"t2",2000);
        startThread(obj,"t3",3000);
        
        
        while(true) {
            try {
                Thread.sleep(1000);
                System.out.println("still alive");
            } catch (InterruptedException e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }
        }
    }
    
    static class TheObject{
        public void doIt(final String message, long timeToSleep) {
            while(true) {
                System.out.println("doit called:"+message);
                try {
                    Thread.sleep(timeToSleep);
                } catch (InterruptedException e) {
                    // TODO Auto-generated catch block
                    e.printStackTrace();
                }
            }
            
        }
    }
    
    private static void startThread(TheObject obj, String message, long timeToSleep) {
        Thread t = new Thread(()->{
            obj.doIt(message,timeToSleep);
        });
        t.setName("Thread-"+message);
        t.start();
    }
}
