package com.ynov.bomberman.stage;

import javafx.scene.image.Image;
import javafx.scene.paint.ImagePattern;
import javafx.scene.shape.Rectangle;

public class Tile {
	public Rectangle tile;

	public boolean isBreakable;
	public boolean isWalkable;

	public int pos;

	public Tile(int x, int y, String tileID, int pos) {
		this.tile = new Rectangle(x * 32, y * 32 + 50, 32, 32);
		
		this.setStyle(tileID);
		
		this.pos = pos;
	}
	
	public void setStyle(String tileID) {
		switch (tileID) {
		case "0":
			Image wall = new Image("/Wall.png");
			tile.setFill(new ImagePattern(wall));
			this.isBreakable = false;
			this.isWalkable = false;
			break;
		case "1":
			Image grass = new Image("/Grass.png");
			tile.setFill(new ImagePattern(grass));
			this.isBreakable = false;
			this.isWalkable = true;
			break;
		case "2":
			Image brick = new Image("/Brick.png");
			tile.setFill(new ImagePattern(brick));
			this.isBreakable = true;
			this.isWalkable = false;
			break;
		default:
			this.isBreakable = false;
			break;
		}
	}
}
