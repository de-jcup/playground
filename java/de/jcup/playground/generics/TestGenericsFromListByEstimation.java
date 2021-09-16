package de.jcup.playground.generics;

import java.lang.reflect.Field;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;

// Startup for answering https://stackoverflow.com/questions/69204360/how-to-get-the-complete-class-type-of-field-in-java
// until it became clear that the question was not correct and it was only about Jackson Typing problems, which can be done by using Jackson correct way
// But here to keep as an example. Maybe a good start point for lib for runtime inspection


/*

Output is:

Variant1: estimation not possible because no content:
>Found type:interface java.util.List
>Estimated content class:class java.lang.Object

Variant2: estimation possible because content available:
>Found type:interface java.util.List
>Estimated content class:class de.jcup.playground.generics.X

*/

class X {
    private String y;
}

class TestGenericsFromListByEstimation {
    
    private List<X> w = new ArrayList<>();
    
    private void run() throws Exception{
        Field f = getClass().getDeclaredField("w");
        Class<?> t = f.getType();// t is "Class of interface java.util.List", but what I want is "Class of java.util.List<X>", how to achieve it
        
        // here we have the only java.util.List<Object> as it is in the byte code!
        // the generic type is not available from code
        System.out.println(">Found type:"+t);
        
        // but at runtime at least an estimation of content is possible:
        Class<?> estimatedContentClass = Object.class;
        Object obj = f.get(this);
        
        if (obj!= null && Iterable.class.isAssignableFrom(t)) {
            Iterable<?> iterable = (Iterable<?>) obj;
            Iterator<?> it = iterable.iterator();
            if (it.hasNext()) {
               /* a very simple approach - just use the first class found here.
                * For a library implementation etc. Every element must be inspected
                * and compared with other found types, than the common interface or class
                * must be found and resolved... 
                * 
                * But here only the simple variant:
                */
                Object element = it.next();
                estimatedContentClass = element.getClass();
            }
        }
        
        System.out.println(">Estimated content class:"+estimatedContentClass);
    }
    
    public static void main(String[] args) throws Exception {
        TestGenericsFromListByEstimation testObject = new TestGenericsFromListByEstimation();
        System.out.println("Variant1: estimation not possible because no content:");
        testObject.run();

        System.out.println("Variant2: estimation possible because content available:");
        testObject.w.add(new X());
        testObject.run(); 
    }
}
