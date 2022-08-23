module com.ynov.bomberman {
    requires javafx.controls;
    requires javafx.fxml;
	requires javafx.graphics;


    opens com.ynov.bomberman to javafx.fxml;
    exports com.ynov.bomberman;
}