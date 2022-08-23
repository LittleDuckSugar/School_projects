package com.ynov.bomberman.player;

import java.util.ArrayList;
import java.util.Iterator;
import java.util.Timer;
import java.util.TimerTask;

import com.ynov.bomberman.stage.Tile;

import javafx.geometry.Rectangle2D;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.scene.layout.Pane;
import javafx.scene.paint.ImagePattern;
import javafx.scene.shape.Circle;
import javafx.scene.shape.Rectangle;
import javafx.util.Duration;

public class Character extends Pane {

//	Configuration du personnage du joueur
	ImageView imageView;
	int count = 3;
	int columns = 3;
	int offSetX = 0;
	int offSetY = 0;
	int width = 32;
	int height = 32;
	public SpriteHandler charachterAnimation;
	
	public int pos = 24;
	public int bombPos = 24;

	public Circle bomb;
	public boolean bombPlanted = false;
	public boolean bombExplosed = false;
	public Timer timerBomb;
	
	public Integer score = 0;
	public boolean win = false;

//	Character est le constructeur du personnage
	public Character(ImageView imageView) {
		this.imageView = imageView;
		this.moveX(32);
		this.moveY(82);
		this.imageView.setViewport(new Rectangle2D(offSetX, offSetY, width, height));
		charachterAnimation = new SpriteHandler(imageView, Duration.millis(200), count, columns, offSetX, offSetY,
				width, height);

		getChildren().addAll(imageView);
	}

//	moveX permet le mouvement du personnage sur l'axe X
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

//	moveY permet le mouvement du personnage sur l'axe Y
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

//	generateBombe permet de générer une bombe sur la tuile du personnage
	public Circle generateBomb(Tile[] mapPlaces) {
		this.bombPlanted = true;

		this.timerBomb = new Timer();

		TimerTask task = new TimerTask() {
			@Override
			public void run() {
				bombPlanted = false;
				bombExplosed = true;

				cancel();
			}
		};
		this.timerBomb.schedule(task, 1500L);

		for (Tile tiles : mapPlaces) {
			if (tiles.tile.intersects(this.getBoundsInParent().getCenterX(),
					this.getBoundsInParent().getCenterY() + 16, width, height)) {
				this.bomb = new Circle(tiles.tile.getX() + 32 / 2, tiles.tile.getY() + 32 / 2, 10,
						new ImagePattern(new Image("/Bomb.png")));
				this.bombPos = tiles.pos;
				break;
			}
		}

//		for (int i = 0; i < mapPlaces.length -1; i++) {
//			if (mapPlaces[i].tile.intersects(this.getBoundsInParent().getCenterX() - 32,
//					this.getBoundsInParent().getCenterY() - 32, width, height)) {
//				this.bomb = new Circle(mapPlaces[i].tile.getX() + 32 / 2, mapPlaces[i].tile.getY() + 32 / 2, 10,
//						new ImagePattern(new Image("/Bomb.png")));
//			}
//		}

		return this.bomb;
	}

}
