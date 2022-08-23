import java.util.Random;
import java.util.Scanner;

public class Jouer {

	public String mode1(int x1, int x2, int x3, int x4) {
		Scanner sc = new Scanner(System.in);		

		String choiceTest;
		int times = 10;

		do {
			StringBuilder sb = new StringBuilder();

			System.out.println("Il vous reste " + times + " essais pour deviner le nombre généré");

			do {
				System.out.println("Entrer un nombre a 4 chiffres");
				choiceTest = sc.next("[0-9]+");

			} while (choiceTest.length() != 4);

			if (x1 == Integer.parseInt(choiceTest) / 1000) {
				sb.append("=");
			} else if (x1 > Integer.parseInt(choiceTest) / 1000) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			if (x2 == Integer.parseInt(choiceTest) / 100 % 10) {
				sb.append("=");
			} else if (x2 > Integer.parseInt(choiceTest) / 100 % 10) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			if (x3 == Integer.parseInt(choiceTest) / 10 % 10) {
				sb.append("=");
			} else if (x3 > Integer.parseInt(choiceTest) / 10 % 10) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			if (x4 == Integer.parseInt(choiceTest) % 10) {
				sb.append("=");
			} else if (x4 > Integer.parseInt(choiceTest) % 10) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			System.out.println(sb.toString());

			if (sb.toString().equals("====")) {
				return "1";
			}

			times = times - 1;
		} while (times != 0);
		return (String.valueOf(x1) + String.valueOf(x2) + String.valueOf(x3) + String.valueOf(x4));
	}

	public String mode2(String choiceGuess) {

		int x1 = Integer.parseInt(choiceGuess.substring(0, 1));
		int x2 = Integer.parseInt(choiceGuess.substring(1, 2));
		int x3 = Integer.parseInt(choiceGuess.substring(2, 3));
		int x4 = Integer.parseInt(choiceGuess.substring(3, 4));

		int test1;
		int test2;
		int test3;
		int test4;

		int choiceTest = 5555;
		String choiceTestdisp;
		int times = 10;

		do {
			StringBuilder sb = new StringBuilder();

			System.out.println("Il reste " + times + " essais pour deviner le nombre choisi");

			if (x1 == (choiceTest / 1000)) {
				sb.append("=");
			} else if (x1 > (choiceTest / 1000)) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			if (x2 == (choiceTest / 100 % 10)) {
				sb.append("=");
			} else if (x2 > (choiceTest / 100 % 10)) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			if (x3 == (choiceTest / 10 % 10)) {
				sb.append("=");
			} else if (x3 > (choiceTest / 10 % 10)) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			if (x4 == (choiceTest % 10)) {
				sb.append("=");
			} else if (x4 > (choiceTest % 10)) {
				sb.append("+");
			} else {
				sb.append("-");
			}

			System.out.println(sb.toString());

			if (sb.toString().equals("====")) {
				return "1";
			}

			if (sb.toString().substring(0, 1).equals("+")) {
				test1 = (choiceTest / 1000) + 1;
			} else if (sb.toString().substring(0, 1).equals("-")) {
				test1 = (choiceTest / 1000) - 1;
			} else {
				test1 = (choiceTest / 1000);
			}

			if (sb.toString().substring(1, 2).equals("+")) {
				test2 = (choiceTest / 100 % 10) + 1;
			} else if (sb.toString().substring(1, 2).equals("-")) {
				test2 = (choiceTest / 100 % 10) - 1;
			} else {
				test2 = (choiceTest / 100 % 10);
			}

			if (sb.toString().substring(2, 3).equals("+")) {
				test3 = (choiceTest / 10 % 10) + 1;
			} else if (sb.toString().substring(2, 3).equals("-")) {
				test3 = (choiceTest / 10 % 10) - 1;
			} else {
				test3 = (choiceTest / 10 % 10);
			}

			if (sb.toString().substring(3, 4).equals("+")) {
				test4 = (choiceTest % 10) + 1;
			} else if (sb.toString().substring(3, 4).equals("-")) {
				test4 = (choiceTest % 10) - 1;
			} else {
				test4 = (choiceTest % 10);
			}

			choiceTestdisp = (
					String.valueOf(test1) + String.valueOf(test2) + String.valueOf(test3) + String.valueOf(test4));
			System.out.println(choiceTestdisp);
			
			choiceTest= Integer.parseInt(choiceTestdisp);

			times = times - 1;
		} while (times != 0);
		return (String.valueOf(x1) + String.valueOf(x2) + String.valueOf(x3) + String.valueOf(x4));
	}

}
