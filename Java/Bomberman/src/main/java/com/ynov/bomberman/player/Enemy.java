package com.ynov.bomberman.player;


import javafx.geometry.Rectangle2D;
import javafx.scene.image.ImageView;
import javafx.scene.layout.Pane;
import javafx.util.Duration;

public class Enemy extends Pane{

	ImageView imageView;
	int height = 32;
	int width = 32;
	
	int offSetX = 32;
	int offSetY = 32;
	
	int count = 3;
	int columns = 3;
	
	public int pos;
	
	public SpriteHandler enemyAnimation;
	
	public Enemy (ImageView imageView) {
		this.imageView = imageView;
		
		this.moveX(64);
		this.moveY(402);
		this.imageView.setViewport(new Rectangle2D(offSetX, offSetY, width, height));
		enemyAnimation = new SpriteHandler(imageView, Duration.millis(200), count, columns, offSetX, offSetY, width, height);
		
		getChildren().addAll(imageView);
	}
	
	public void moveX(int x) {
		boolean right = x > 0 ? true : false;
		for (int i = 0; i < Math.abs(x); i++) {
			if (right) {
				this.setTranslateX(this.getTranslateX() + 1);
			} else {
				this.setTranslateX(this.getTranslateX() - 1);
			}
		}
	}

	public void moveY(int y) {
		boolean right = y > 0 ? true : false;
		for (int i = 0; i < Math.abs(y); i++) {
			if (right) {
				this.setTranslateY(this.getTranslateY() + 1);
			} else {
				this.setTranslateY(this.getTranslateY() - 1);
			}
		}
	}
}
