package com.ynov.bomberman.menu;

import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.geometry.Pos;
import javafx.scene.control.Button;
import javafx.scene.layout.StackPane;
import javafx.scene.paint.Color;
import javafx.scene.paint.CycleMethod;
import javafx.scene.paint.LinearGradient;
import javafx.scene.paint.Stop;
import javafx.scene.text.Font;
import javafx.scene.text.FontWeight;

public class MenuItem extends StackPane {
	public Button button;
    public MenuItem(String name) {
        LinearGradient gradient = new LinearGradient(0, 0, 1, 0, true, CycleMethod.NO_CYCLE, new Stop[] {
                new Stop(0, Color.DARKBLUE),
                new Stop(0.1, Color.BLACK),
                new Stop(0.9, Color.BLACK),
                new Stop(1, Color.DARKBLUE)

        });


        this.button = new Button(name);
        button.setTextFill(Color.BLACK);
        button.setFont(Font.font("Times New Roman", FontWeight.SEMI_BOLD,20));

        setAlignment(Pos.CENTER);
        getChildren().addAll(button);
        setOnMouseEntered(event -> {
            button.setTextFill(Color.WHITE);

        });

        setOnMouseExited(event -> {
            button.setTextFill(Color.BLACK);
        });

    }
    public void setOnAction(EventHandler<ActionEvent> actionEventEventHandler) {
    }
}

