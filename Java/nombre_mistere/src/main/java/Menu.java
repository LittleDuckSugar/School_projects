import java.util.Random;
import java.util.Scanner;

public class Menu {

	public static void main(String[] args) {
		System.out.println("Bienvenu dans Nombre mistere...");

		Jouer jeu = new Jouer();

		int choice = 0;
		String res;

		do {
			System.out.println("1. Deviner le nombre");
			System.out.println("2. Faire deviner le nombre");
			System.out.println("3. Quitter l'application");

			Scanner sc = new Scanner(System.in);
			choice = sc.nextInt();

			switch (choice) {
			case 1:
				Random r = new Random();
				int x1 = r.nextInt(9);
				int x2 = r.nextInt(9);
				int x3 = r.nextInt(9);
				int x4 = r.nextInt(9);
				res = jeu.mode1(x1, x2, x3, x4);
				if (res == "1") {
					System.out.println("Bravo vous avez trouvé le bon nombre!");
				} else {
					System.out.println("Dommage, vous avez perdu, vous deviez trouver " + res);
				}
				break;
			case 2:
				System.out.println("Entrer ce que doit trouver l'ordinateur ?");
				String choiceGuess = sc.next("[0-9]+");

				res = jeu.mode2(choiceGuess);
				if (res == "1") {
					System.out.println("L'ordinateur a trouvé le bon nombre!");
				} else {
					System.out.println("Dommage, l'ordinateur a perdu, il devait trouver " + res);
				}
				break;
			case 3:
				break;
			default:
				System.out.println("Merci de saisir une valeur entre 1 et 3.");
				break;
			}

		} while (choice != 3);

		System.out.println("Fin du programme");

	}

}
