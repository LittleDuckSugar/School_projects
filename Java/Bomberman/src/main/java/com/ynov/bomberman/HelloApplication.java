package com.ynov.bomberman;

import javafx.animation.AnimationTimer;
import javafx.application.Application;
import javafx.event.ActionEvent;
import javafx.event.EventHandler;
import javafx.scene.Scene;
import javafx.scene.canvas.Canvas;
import javafx.scene.canvas.GraphicsContext;
import javafx.scene.control.Button;
import javafx.scene.control.TableColumn;
import javafx.scene.control.TableView;
import javafx.scene.control.TextArea;
import javafx.scene.control.cell.MapValueFactory;
import javafx.scene.effect.InnerShadow;
import javafx.scene.image.Image;
import javafx.scene.image.ImageView;
import javafx.scene.input.KeyCode;
import javafx.scene.layout.Pane;
import javafx.scene.paint.Color;
import javafx.scene.text.Font;
import javafx.scene.text.FontPosture;
import javafx.scene.text.FontWeight;
import javafx.scene.text.Text;
import javafx.stage.Stage;

import java.io.IOException;
import java.util.ArrayList;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Timer;
import java.util.TimerTask;

import javafx.scene.Group;
import javafx.animation.Timeline;

import com.ynov.bomberman.menu.Menu;
import com.ynov.bomberman.menu.MenuItem;
import com.ynov.bomberman.menu.Title;
import com.ynov.bomberman.player.Character;
import com.ynov.bomberman.player.Enemy;
import com.ynov.bomberman.stage.Game;
import com.ynov.bomberman.stage.Tile;

public class HelloApplication extends Application {

	public static final int WIDTH = 736;
	public static final int HEIGHT = 466;
	TextArea inputName;
	Stage stage;
	List<Scene> listScenes;
	List<Group> listGroups;
	Group groupMenuPage;
	Group groupGamePage;
	Group groupScorePage;
	Scene sceneMenu;
	Scene sceneGameInitial;
	Scene sceneScore;
	Text score;
	String namePlayer;
	TableView tableScore;
	ArrayList<Map<String, Object>> scores;
	Timeline tl;
	Button buttonReturnMenu;
	Button buttonToHighScore;
	Group group = new Group();
	Group group1 = new Group();
	int indexActiveSceneGame = 0;
	private HashMap<KeyCode, Boolean> keys = new HashMap<>();
	boolean isMouvement = false;
	Text text = new Text();
	static boolean gameOver = false;
//	Initialisation du joueur
	Character playerOne = new Character(new ImageView(new Image("/RPGMaker.png")));

//	Initialisation des ennemies
	Enemy onil = new Enemy(new ImageView(new Image("/Ballom.png")));

	static Pane root = new Pane();

	@Override
	public void start(Stage stage) throws IOException {
		listScenes = new ArrayList<Scene>();
		listGroups = new ArrayList<Group>();

		// MenuPage
		groupMenuPage = new Group();
		ImageView img = new ImageView(new Image("/bgMenu.png"));
		img.setFitWidth(WIDTH);
		img.setFitHeight(HEIGHT);
		groupMenuPage.getChildren().add(img);

		Title title = new Title("Bomberman");
		title.setTranslateX(50);
		title.setTranslateY(200);
		MenuItem startGame = new MenuItem("NEW GAME");
		MenuItem highScore = new MenuItem("HIGHSCORE");
		MenuItem exit = new MenuItem("EXIT");
		Menu vbox = new Menu(startGame, highScore, exit);
		vbox.setTranslateX(100);
		vbox.setTranslateY(300);
		inputName = new TextArea();
		inputName.setPrefHeight(30);
		inputName.setPrefWidth(200);
		inputName.setTranslateX(100);
		inputName.setTranslateY(420);
		groupMenuPage.getChildren().addAll(title, vbox, inputName);

		sceneMenu = new Scene(groupMenuPage, WIDTH, HEIGHT);
		stage.setTitle("BOMBERMAN");
		stage.setScene(sceneMenu);
		
		// Game
		groupGamePage = initGame();
		listGroups.add(groupGamePage);
		sceneGameInitial = new Scene(groupGamePage, WIDTH, HEIGHT, Color.GREY);
		sceneGameInitial.setOnKeyPressed(event -> keys.put(event.getCode(), true));
		sceneGameInitial.setOnKeyReleased(event -> keys.put(event.getCode(), false));
		listScenes.add(sceneGameInitial);
		
		//Score
		groupScorePage= new Group();
		groupScorePage=initScore();
		sceneScore=new Scene(groupScorePage,WIDTH,HEIGHT,Color.GRAY);
		stage.show();
		startGame.button.setOnAction((EventHandler<ActionEvent>) new EventHandler<ActionEvent>() {
			@Override
			public void handle(ActionEvent event) {
				stage.setScene(sceneGameInitial);
				stage.show();
			}
		});
		highScore.button.setOnAction((EventHandler<ActionEvent>) new EventHandler<ActionEvent>() {
			@Override
			public void handle(ActionEvent event) {
				stage.setScene(sceneScore);
				stage.show();
			}
		});
		buttonReturnMenu.setOnAction((EventHandler<ActionEvent>) new EventHandler<ActionEvent>() {
			@Override
			public void handle(ActionEvent event) {
				stage.setScene(sceneMenu);
				stage.show();
			}
		});
		exit.button.setOnAction((EventHandler<ActionEvent>) new EventHandler<ActionEvent>() {
			@Override
			public void handle(ActionEvent event) {
				stage.close();
			}
		});
		
	}

	private Group initGame() {
//		Sauvegarde de toute les tuiles de la carte
		Tile[] mapPlaces = new Tile[299];

		int y = 0;
		int pos = 0;

		for (String line : Game.LEVEL1) {
			int x = 0;
			for (String type : line.split("")) {
				Tile bloc = new Tile(x, y, type, pos);

				mapPlaces[pos] = bloc;

				group.getChildren().add(bloc.tile);
				x++;
				pos++;
			}
			y++;
		}

		group.getChildren().add(playerOne);
		group.getChildren().add(onil);

		
//      Score
		
		text.setText("Score " + playerOne.score.toString());
		text.setX(25);
		text.setY(32);
		group.getChildren().add(text);
		text.setFill(Color.WHITE);
		text.setFont(Font.font("Verdana", 25));
		
		Canvas c = new Canvas(WIDTH,HEIGHT);
		GraphicsContext gc = c.getGraphicsContext2D();
		group.getChildren().add(c);
//		Actions en boucle
		AnimationTimer timer = new AnimationTimer() {
			long lastTick = 0;
			private Button buttonToHighScore;
			@Override
			public void handle(long now) {
				if (!playerOne.win) {
					characterMovement(mapPlaces);
					if (!isMouvement){
						enemyMovement(mapPlaces);
					}
					DeadHandler();
					bombHandler(mapPlaces);
				}
				if (lastTick == 0 && gameOver) {
					lastTick = now;
					addScoreToTable();
					this.buttonToHighScore= new Button("HighScore");
					buttonToHighScore.setLayoutX(280);
					buttonToHighScore.setLayoutY(300);
					buttonToHighScore.setFont(Font.font("Arial", FontWeight.EXTRA_BOLD, FontPosture.ITALIC, 30));
					buttonToHighScore.setEffect(new InnerShadow(10, Color.DARKRED));
					group.getChildren().add(buttonToHighScore);
					group.getChildren().remove(onil);
					group.getChildren().remove(playerOne);
					this.buttonToHighScore.setOnAction((EventHandler<ActionEvent>) new EventHandler<ActionEvent>() {
						@Override
						public void handle(ActionEvent event) {
							stage.setScene(sceneScore);
							stage.show();
						}
					});

					tick(gc);
					return;
				}

				
			}
		};
		timer.start();
		return group;
	}
	private Group initScore(){
		tableScore= new TableView();
		scores = new ArrayList<Map<String, Object>>();
		TableColumn<Map, String> col1 = new TableColumn<>("Name");
        col1.setCellValueFactory(new MapValueFactory<>("name"));

        TableColumn<Map, String> col2 = new TableColumn<>("Score");
        col2.setCellValueFactory(new MapValueFactory<>("score"));
        
        for (Map<String, Object> item:scores) {
            tableScore.getItems().addAll(item);
        }
        tableScore.getColumns().add(col1);
        tableScore.getColumns().add(col2);
        tableScore.setColumnResizePolicy(TableView.CONSTRAINED_RESIZE_POLICY);
        tableScore.setLayoutX(WIDTH/2);
        tableScore.setLayoutY(HEIGHT/8);
        group1.getChildren().add(tableScore);
        buttonReturnMenu= new Button("Return to menu");
        group1.getChildren().add(buttonReturnMenu);
		return group1;
	}

//	characterMovement prend en charge les mouvements du joueur
	public void characterMovement(Tile[] mapPlaces) {

		if (isPress(KeyCode.Z)) {
			for (int i = 0; i < mapPlaces.length; i++) {
				if ((playerOne.getBoundsInParent().getCenterX() >= mapPlaces[i].tile.getX()
						&& playerOne.getBoundsInParent().getCenterX() <= mapPlaces[i].tile.getX() + 32)
						&& (playerOne.getBoundsInParent().getCenterY() - 2 + 16 >= mapPlaces[i].tile.getY()
								&& playerOne.getBoundsInParent().getCenterY() - 2 + 16 <= mapPlaces[i].tile.getY()
										+ 32)) {
					//System.out.println("Player is on case " + mapPlaces[i].pos + "from math");
					//System.out.println("Player is on case " + playerOne.pos + "from player infos");
					if (mapPlaces[i].isWalkable) {
						playerOne.charachterAnimation.play();
						playerOne.charachterAnimation.setOffsetY(96);
						playerOne.moveY(-2);
						playerOne.pos = mapPlaces[i].pos;
					} else {
						playerOne.charachterAnimation.stop();
					}
					break;
				}
			}

		} else if (isPress(KeyCode.S)) {

			for (int i = 0; i < mapPlaces.length; i++) {
				if ((playerOne.getBoundsInParent().getCenterX() >= mapPlaces[i].tile.getX()
						&& playerOne.getBoundsInParent().getCenterX() <= mapPlaces[i].tile.getX() + 32)
						&& (playerOne.getBoundsInParent().getCenterY() + 2 + 16 >= mapPlaces[i].tile.getY()
								&& playerOne.getBoundsInParent().getCenterY() + 2 + 16 <= mapPlaces[i].tile.getY()
										+ 32)) {
					//System.out.println("Player is on case " + mapPlaces[i].pos + "from math");
					//System.out.println("Player is on case " + playerOne.pos + "from player infos");
					if (mapPlaces[i].isWalkable) {
						playerOne.charachterAnimation.play();
						playerOne.charachterAnimation.setOffsetY(0);
						playerOne.moveY(2);
						playerOne.pos = mapPlaces[i].pos;
					} else {
						playerOne.charachterAnimation.stop();
					}
					break;
				}
			}
		} else if (isPress(KeyCode.D)) {

			for (int i = 0; i < mapPlaces.length; i++) {

				if ((playerOne.getBoundsInParent().getCenterX() + 2 >= mapPlaces[i].tile.getX()
						&& playerOne.getBoundsInParent().getCenterX() + 2 <= mapPlaces[i].tile.getX() + 32)
						&& (playerOne.getBoundsInParent().getCenterY() + 16 >= mapPlaces[i].tile.getY()
								&& playerOne.getBoundsInParent().getCenterY() + 16 <= mapPlaces[i].tile.getY() + 32)) {
					//System.out.println("Player is on case " + mapPlaces[i].pos + "from math");
					//System.out.println("Player is on case " + playerOne.pos + "from player infos");
					if (mapPlaces[i].isWalkable) {
						playerOne.charachterAnimation.play();
						playerOne.charachterAnimation.setOffsetY(64);
						playerOne.moveX(2);
						playerOne.pos = mapPlaces[i].pos;
					} else {
						playerOne.charachterAnimation.stop();
					}
					break;
				}
			}

		} else if (isPress(KeyCode.Q)) {
			for (int i = 0; i < mapPlaces.length; i++) {
				if ((playerOne.getBoundsInParent().getCenterX() - 2 >= mapPlaces[i].tile.getX()
						&& playerOne.getBoundsInParent().getCenterX() - 2 <= mapPlaces[i].tile.getX() + 32)
						&& (playerOne.getBoundsInParent().getCenterY() + 16 >= mapPlaces[i].tile.getY()
								&& playerOne.getBoundsInParent().getCenterY() + 16 <= mapPlaces[i].tile.getY() + 32)) {
					//System.out.println("Player is on case " + mapPlaces[i].pos + "from math");
					//System.out.println("Player is on case " + playerOne.pos + "from player infos");
					if (mapPlaces[i].isWalkable) {
						playerOne.charachterAnimation.play();
						playerOne.charachterAnimation.setOffsetY(32);
						playerOne.pos = mapPlaces[i].pos;
						playerOne.moveX(-2);
					} else {
						playerOne.charachterAnimation.stop();
					}
					break;
				}
			}
		} else {
			playerOne.charachterAnimation.stop();
		}
	}

//	enemyMovement prend en charge les mouvements des ennemies
	public void enemyMovement(Tile[] mapPlaces) {

		for (int i = 0; i < mapPlaces.length; i++) {
			if ((onil.getBoundsInParent().getCenterX() >= mapPlaces[i].tile.getX()
					&& onil.getBoundsInParent().getCenterX() <= mapPlaces[i].tile.getX() + 32)
					&& (onil.getBoundsInParent().getCenterY() + 16 >= mapPlaces[i].tile.getY()
							&& onil.getBoundsInParent().getCenterY() + 16 <= mapPlaces[i].tile.getY() + 32)) {
				
				// System.out.println(mapPlaces[i].pos);
				onil.pos = mapPlaces[i].pos;
				
				ArrayList<Integer> mouvementAllow = new ArrayList<>();
				if (mapPlaces[i - 1].isWalkable) {
					mouvementAllow.add(- 1);
				}
				if (mapPlaces[i + 1].isWalkable) {
					mouvementAllow.add(+ 1);
				}
				if (mapPlaces[i + 23].isWalkable) {
					mouvementAllow.add(+ 23);
				}
				if (mapPlaces[i - 23].isWalkable) {
					mouvementAllow.add(- 23);
				}
				
				int random = (int)(Math.random()*(mouvementAllow.size()));
				int mouvementToDo = mouvementAllow.get(random);
				
				if (mouvementToDo == 23 ){
					Timer t = new Timer();
				    TimerTask task = new TimerTask() {
				      int i=0;
				      public void run() {
				    	  onil.moveY(2);
				        if(i == 32) {
				        	isMouvement = false;
				        	t.cancel();
				        }
				        i += 2;
				      }
				    };
				    
				    isMouvement = true;
				    t.schedule(task, new Date(), 50);

				}
				if (mouvementToDo == - 23 ){
					Timer t = new Timer();
				    TimerTask task = new TimerTask() {
				      int i=0;
				      public void run() {
				    	  onil.moveY(- 2);
				        if(i == 32) {
				        	isMouvement = false;
				        	t.cancel();
				        }
				        i += 2;
				      }
				    };
				    
				    isMouvement = true;
				    t.schedule(task, new Date(), 50);
				}
				if (mouvementToDo == 1 ){
					Timer t = new Timer();
				    TimerTask task = new TimerTask() {
				      int i=0;
				      public void run() {
				    	  onil.moveX(2);
				        if(i == 32) {
				        	isMouvement = false;
				        	t.cancel();
				        }
				        i += 2;
				      }
				    };
				    
				    isMouvement = true;
				    t.schedule(task, new Date(), 50);
					
				}
				if (mouvementToDo == - 1 ){
					Timer t = new Timer();
				    TimerTask task = new TimerTask() {
				      int i=0;
				      public void run() {
				    	  onil.moveX(- 2);
				        if(i == 32) {
				        	isMouvement = false;
				        	t.cancel();
				        }
				        i += 2;
				      }
				    };
				    
				    isMouvement = true;
				    t.schedule(task, new Date(), 50);
			}	
				break;
				}
		}
	}
	public void tick(GraphicsContext gc) {
		if (gameOver) {
			gc.setFill(Color.RED);
			gc.setFont(Font.font("Arial", FontWeight.EXTRA_BOLD, FontPosture.ITALIC, 55));
			gc.setEffect(new InnerShadow(10, Color.DARKRED));
			gc.fillText("GAME OVER", 200, 250);
			
			return;
		}
	}
	public void DeadHandler() {
		
		if ((onil.getBoundsInParent().getCenterX() >= playerOne.getBoundsInParent().getCenterX()
                && onil.getBoundsInParent().getCenterX() <= playerOne.getBoundsInParent().getCenterX() + 32)
                && (onil.getBoundsInParent().getCenterY() >= playerOne.getBoundsInParent().getCenterY()
                        && onil.getBoundsInParent().getCenterY() <= playerOne.getBoundsInParent().getCenterY()+ 32)) {
			gameOver = true;
		}
	}

    
//	bombHandler supporte la pose et l'explosion des bombes du joueur
	public void bombHandler(Tile[] mapPlaces) {
		this.stage= new Stage();
		if (playerOne.bombExplosed) {

//			multiple de 23
//			-----------------------
//			---------b*b-----b-----
//			----------b-----b*b----
//			-----------------b-----
//			si * est 44eme element alors : 
//			 - 44 - 23 = explosion en haut
//			 - 44 - 1 = explotion à gauche
//			 - 44 + 1 = explotion à droite
//			 - 44 + 23 = explosion en dessous

			for (int i = 0; i < mapPlaces.length; i++) {
				if (playerOne.bomb.getCenterX() - 16 == mapPlaces[i].tile.getX()
						&& playerOne.bomb.getCenterY() - 16 == mapPlaces[i].tile.getY()) {

					if (mapPlaces[i + 1].isBreakable) {
						mapPlaces[i + 1].setStyle("1");
						playerOne.score += 100;
						text.setText("Score " + playerOne.score.toString());
					}

					if (mapPlaces[i - 1].isBreakable) {
						mapPlaces[i - 1].setStyle("1");
						playerOne.score += 100;
						text.setText("Score " + playerOne.score.toString());
					}

					if (mapPlaces[i + 23].isBreakable) {
						mapPlaces[i + 23].setStyle("1");
						playerOne.score += 100;
						text.setText("Score " + playerOne.score.toString());
					}

					if (mapPlaces[i - 23].isBreakable) {
						mapPlaces[i - 23].setStyle("1");
						playerOne.score += 100;
						text.setText("Score " + playerOne.score.toString());
					}

					if (playerOne.pos == mapPlaces[i].pos || playerOne.pos == mapPlaces[i + 1].pos
							|| playerOne.pos == mapPlaces[i - 1].pos || playerOne.pos == mapPlaces[i + 23].pos
							|| playerOne.pos == mapPlaces[i - 23].pos) {
//						Handle death here
						group.getChildren().remove(playerOne);
						group.getChildren().remove(onil);
						addScoreToTable();
						System.out.println(playerOne.score);
						stage.setScene(sceneScore);
						stage.show();
					}
					
					if (onil.pos == mapPlaces[i].pos || onil.pos == mapPlaces[i + 1].pos
							|| onil.pos == mapPlaces[i - 1].pos || onil.pos == mapPlaces[i + 23].pos
							|| onil.pos == mapPlaces[i - 23].pos) {
						
						group.getChildren().remove(onil);
						group.getChildren().remove(playerOne);
						addScoreToTable();
						playerOne.win = true;
						stage.setScene(sceneScore);
						stage.show();
						
					}

					playerOne.toFront();
				}
			}

			group.getChildren().remove(playerOne.bomb);

			playerOne.bombExplosed = false;
		}

		if (!playerOne.bombPlanted && isPress(KeyCode.SPACE)) {
			group.getChildren().add(playerOne.generateBomb(mapPlaces));
		}
	}

	private boolean isPress(KeyCode key) {
		return keys.getOrDefault(key, false);
	}
	private void addScoreToTable() {
        namePlayer = inputName.getText();
        HashMap<String, Object> toAdd = new HashMap<>();
        toAdd.put("userName", namePlayer);
        toAdd.put("score", playerOne.score);
        scores.add(toAdd);
        tableScore.getItems().add(toAdd);
    }

	public static void main(String[] args) {
		launch();
	}
}