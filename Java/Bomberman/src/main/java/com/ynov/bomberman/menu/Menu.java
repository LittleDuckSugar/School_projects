package com.ynov.bomberman.menu;

import javafx.scene.layout.VBox;
import javafx.scene.paint.Color;
import javafx.scene.shape.Line;

public class Menu extends VBox {
    public Menu(MenuItem...items) {
        getChildren().add(createSeperator());
        for(MenuItem item : items) {
            getChildren().addAll(item, createSeperator());
        }
    }

    private Line createSeperator() {
        Line sep = new Line();
        sep.setEndX(210);
        sep.setStroke(Color.DARKGREY);
        return sep;
    }
}
