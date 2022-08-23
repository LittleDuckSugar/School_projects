package com.ynov.bomberman.player;

import javafx.animation.Animation;
import javafx.animation.Interpolator;
import javafx.animation.Transition;
import javafx.geometry.Rectangle2D;
import javafx.scene.image.ImageView;
import javafx.util.Duration;

public class SpriteHandler extends Transition {

	private final ImageView imageView;
	private final int count;
	private final int columns;
	private int offsetX;
	private int offsetY;
	private final int width;
	private final int height;

//	SpriteHandler est le constructeur de la class
	public SpriteHandler(ImageView imageView, Duration duration, int count, int columns, int offsetX, int offsetY,
			int width, int height) {
		this.imageView = imageView;
		this.count = count;
		this.columns = columns;
		this.setOffsetX(offsetX);
		this.setOffsetY(offsetY);
		this.width = width;
		this.height = height;

		setCycleDuration(duration);
		setCycleCount(Animation.INDEFINITE);
		setInterpolator(Interpolator.LINEAR);
		this.imageView.setViewport(new Rectangle2D(offsetX, offsetY, width, height));
	}

//	Met a jour le sprite en X
	public void setOffsetX(int offsetX) {
		this.offsetX = offsetX;
	}

//	Met a jour le sprite en Y
	public void setOffsetY(int offsetY) {
		this.offsetY = offsetY;
	}

//	permet de faire le changement d'image du sprite pour chaque déplacement
	protected void interpolate(double frac) {
		final int index = Math.min((int) Math.floor(count * frac), count - 1);
		final int x = (index % columns) * width + offsetX;
		final int y = (index / columns) * height + offsetY;
		imageView.setViewport(new Rectangle2D(x, y, width, height));
	}

}
