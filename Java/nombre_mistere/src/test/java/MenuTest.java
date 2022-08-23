import static org.junit.jupiter.api.Assertions.assertEquals;

import java.io.ByteArrayInputStream;
import java.io.InputStream;

import org.junit.jupiter.api.Test;

public class MenuTest extends Menu {

	@Test
	public void testMainexit() {
		Menu testaccess = new Menu();

		InputStream sysInBackup = System.in; // backup System.in to restore it later
		String simulatedUserInput = "3" + System.getProperty("line.separator");
		ByteArrayInputStream in = new ByteArrayInputStream(simulatedUserInput.getBytes());
		System.setIn(in);

		testaccess.main(null);
		System.setIn(sysInBackup);

	}
}
