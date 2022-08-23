
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.io.ByteArrayInputStream;
import java.io.InputStream;

import org.junit.jupiter.api.Test;

public class JouerTest {

	@Test
	public void testmode1true() {
		// Arrange
		InputStream sysInBackup = System.in; // backup System.in to restore it later
		ByteArrayInputStream in = new ByteArrayInputStream("5555\n".getBytes());
		System.setIn(in);

		Jouer jeu = new Jouer();

		// Act
		String res = jeu.mode1(5, 5, 5, 5);

		// Assert
		assertTrue(res.equals("1"));

	}

	@Test
	public void testmode1trueup() {
		// Arrange
		InputStream sysInBackup = System.in; // backup System.in to restore it later
		String simulatedUserInput = "1111" + System.getProperty("line.separator") + "2222"
				+ System.getProperty("line.separator") + "3333" + System.getProperty("line.separator");
		ByteArrayInputStream in = new ByteArrayInputStream(simulatedUserInput.getBytes());
		System.setIn(in);

		System.out.println("Test de mode1trueup");
		Jouer jeu = new Jouer();

		// Act
		String res = jeu.mode1(3, 3, 3, 3);
		System.setIn(sysInBackup);

		// Assert
		assertTrue(res.equals("1"));

	}

	@Test
	public void testmode1truedown() {
		// Arrange
		InputStream sysInBackup = System.in; // backup System.in to restore it later
		String simulatedUserInput = "7777" + System.getProperty("line.separator") + "6666"
				+ System.getProperty("line.separator") + "5555" + System.getProperty("line.separator");
		ByteArrayInputStream in = new ByteArrayInputStream(simulatedUserInput.getBytes());
		System.setIn(in);

		System.out.println("Test de mode1truedown");
		Jouer jeu = new Jouer();

		// Act
		String res = jeu.mode1(5, 5, 5, 5);
		System.setIn(sysInBackup);

		// Assert
		assertTrue(res.equals("1"));

	}

	@Test
	public void testmode1false() {
		// Arrange
		InputStream sysInBackup = System.in; // backup System.in to restore it later
		String simulatedUserInput = "1111" + System.getProperty("line.separator") + "2222"
				+ System.getProperty("line.separator") + "3333" + System.getProperty("line.separator") + "2222"
				+ System.getProperty("line.separator") + "2222" + System.getProperty("line.separator") + "2222"
				+ System.getProperty("line.separator") + "2222" + System.getProperty("line.separator") + "2222"
				+ System.getProperty("line.separator") + "2222" + System.getProperty("line.separator") + "2222"
				+ System.getProperty("line.separator");
		ByteArrayInputStream in = new ByteArrayInputStream(simulatedUserInput.getBytes());
		System.setIn(in);

		System.out.println("Test de mode1false");
		Jouer jeu = new Jouer();

		// Act
		String res = jeu.mode1(4, 4, 4, 4);
		System.setIn(sysInBackup);

		// Assert
		assertTrue(res.equals("4444"));

	}

	@Test
	public void testmode2under() {
		// Arrange
		Jouer jeu = new Jouer();

		// Act
		String res = jeu.mode2("0341");

		// Assert
		assertTrue(res.equals("1"));

	}

	@Test
	public void testmode2upper() {
		// Arrange
		Jouer jeu = new Jouer();

		// Act
		String res = jeu.mode2("6798");

		// Assert
		assertTrue(res.equals("1"));

	}

}
