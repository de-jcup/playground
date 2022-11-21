package de.jcup.playground.locale;

import java.util.Locale;

public class LocaleToLowercase {

	public static void main(String[] args) {
		String warning = "WARNING";
		Locale trlocale= Locale.forLanguageTag("tr-TR");
		
		String lowercased = warning.toLowerCase(trlocale);
		
		System.out.println("For a TR locale, the lowercase output even for standard ascii parts can change!" );
		System.out.println(warning);
		System.out.println("becomes...");
		System.out.println("-------- with locale tr-TR:------");
		System.out.println(lowercased);
		System.out.println("-------- with locale ROOT:------");
		System.out.println(warning.toLowerCase(Locale.ROOT));
    
	}
}
