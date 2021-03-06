package com.cah;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collection;
import java.util.List;
//import java.util.Random;
import java.util.concurrent.BlockingQueue;
import java.util.concurrent.atomic.AtomicBoolean;

import android.app.AlertDialog;
import android.content.Context;
import android.graphics.Bitmap;
import android.graphics.BitmapFactory;
import android.graphics.Color;
import android.util.Pair;
import android.view.View;
import android.view.View.OnClickListener;
import android.widget.HorizontalScrollView;
import android.widget.LinearLayout;
import android.widget.LinearLayout.LayoutParams;
import android.widget.TextView;

import com.cah.customviews.CardView;
import com.cah.customviews.GameTable;
import com.cah.datastructures.Card;

/**
 * Class that handles all gameplay interactions with the UI
 * The user mainly directly interacts with this class.
 */
public class CahPlayer {

	final AtomicBoolean go = new AtomicBoolean(true);
	final Cah cahActivity;
	final CahClient client;
	String tableID;
	int playerId;
	int numberOfPlayers = 0;
	AtomicBoolean playerIsCzar = new AtomicBoolean(false);
	AlertDialog czarDialog;
	int numberOfRoundsWon = 0;
	ArrayList<Pair<Integer, String>> czarWhiteCards = new ArrayList<Pair<Integer, String>>();
	Thread messageHandler;
	public final static List<Player> currentList = new ArrayList<Player>();

	//mehbe?
	//public final static List<String> imageList = new ArrayList<String>();
	int imageID;

	/**
	 * @param cahActivity
	 * @param client
	 * @param tableToJoin
	 *            Table ID to join. Creates a new table if null
	 */
	public CahPlayer(Cah cahActivity, CahClient client, String tableToJoin) {
		this.cahActivity = cahActivity;
		this.client = client;
		this.tableID = tableToJoin;

		//Random rand = new Random();
		//imageID = rand.nextInt(6);
		
		imageID = playerId;

	    Cah.performOnBackgroundThread(new Runnable() {

			@Override
			public void run() {
				try {
					messageHandler = new Thread(new HandleMessageThread());
					messageHandler.start();

					// Go ahead and join/create table at this point.
					if (CahPlayer.this.tableID != null) {
						// Join table
						CahPlayer.this.client.outgoing.put(new TableDelta(
								"join", CahPlayer.this.tableID));
					} else {
						// Create new table.
						CahPlayer.this.client.outgoing.put(new TableDelta(
								"new", null));
					}
				} catch (InterruptedException e) {
					// TODO: Die gracefully
					e.printStackTrace();
				}
			}
		}); // End performOnBackgroundThread
	}

	//handles incoming messages from server
	public void handleIncomingMessages(BlockingQueue<Delta> incoming, BlockingQueue<Delta> outgoing) throws InterruptedException {
		
		for (int i = 0; i < currentList.size(); i++){
			Player player = currentList.get(i);
			System.out.print("Player ID is" + player.id + ",");
			System.out.println("Player czar status is" + player.czar);
		}
		
		while( go.get() ) {
			// This will block until something comes in.
			Delta incoming_message = incoming.take(); 
			
			// Print out debug information
			System.out.println("in handleIncomingMessages(): " + incoming_message.toString());
			this.showDebugText(incoming_message.toString());

			Class<? extends Delta> c = incoming_message.getClass();
			if(c == TableDelta.class){
				//TODO: Implement this type of delta.
				TableDelta delta = (TableDelta) incoming_message;
				if(delta.Command.equals("ok")) {
					this.tableID = delta.Id;
					//ask for an id
					outgoing.put(new PlayerDelta(0, "my-id?"));
				}
			} else if (c == DeckDelta.class){
				//TODO: Implement this type of delta.
				final DeckDelta delta = (DeckDelta) incoming_message;
				if(delta.DeckTo.equals("hand") && delta.DeckFrom.equals("white-draw")) {
					// Add the card to our hand
					for(Card card : delta.Cards) {
						addCardToHand(new Card(Card.Color.WHITE, card.text));
					}
				} else if (delta.DeckTo.equals("play") && delta.DeckFrom.equals("black-draw") && this.playerIsCzar.get() == true) {
					// Show AlertDialog showing only the black card.
					cahActivity.runOnUiThread(new Runnable() {
						@Override
						public void run() {
							CahPlayer.this.playerIsCzar.set(true);
							AlertDialog.Builder dialogBuilder = new AlertDialog.Builder(cahActivity);
							dialogBuilder.setView(CzarActivity.getBlackCardView(delta.Cards[0], cahActivity));
							dialogBuilder.setCancelable(false);
							czarDialog = dialogBuilder.show();
						}
					});

					//TODO: Dismiss the dialog when we recieve all player's white cards.
				} else if (delta.DeckTo.equals("play") && this.playerIsCzar.get() == true) {
					
					for(Card card : delta.Cards){
						czarWhiteCards.add(new Pair<Integer, String>(delta.Player, card.text));
					}
					
					if(czarWhiteCards.size() == numberOfPlayers-1) {
						// All players have played a card. Czar should now choose best card.
						System.out.println("SHOWING WHITE CARD CHOOSER");
						showWhiteCardChooser(czarWhiteCards, cahActivity);
					}
					
				} else if (delta.DeckTo.equals("winner") && delta.Player == playerId){
					cahActivity.runOnUiThread(new Runnable() {

						@Override
						public void run() {
							numberOfRoundsWon++;
							AlertDialog.Builder dialogBuilder = new AlertDialog.Builder(cahActivity);
							dialogBuilder.setMessage("Congratulations! You won that round! You have won a total of " + numberOfRoundsWon + " rounds.");
							dialogBuilder.setPositiveButton("Ok", null);
							dialogBuilder.show();
						}
						
					});
				}
			} else if (c == PlayerDelta.class) {
				//TODO: Implement this type of delta.
				PlayerDelta delta = (PlayerDelta) incoming_message;
				if(delta.Message.equals("your-id")) {
					this.playerId = delta.Id;
					outgoing.put(new PlayerDelta(this.playerId, "join"));
					playerJoined(this.playerId, false);
					numberOfPlayers++;
				} else if (delta.Id != this.playerId && delta.Message.equals("join")){
					// Another player has joined the table
					playerJoined(delta.Id, false);
					cahActivity.runOnUiThread(new Runnable() {
						@Override
						public void run() {
							cahActivity.playerCanPlayCard(true);
						}
					});
					numberOfPlayers++;
				} else if (delta.Message.equals("leave")) {
					//TODO: Remove player from table here
					playerLeft(delta.Id);
					numberOfPlayers--;
				} else if (delta.Message.equals("is-czar")) {
					cahActivity.runOnUiThread(new Runnable() {

						@Override
						public void run() {
							cahActivity.playerCanPlayCard(true);
						}
						
					});
					if(delta.Id == this.playerId){
						this.playerIsCzar.set(true);
						playerBecomesCzar(delta.Id);
						cahActivity.runOnUiThread(new Runnable() {

							@Override
							public void run() {
								cahActivity.playerCanPlayCard(false);
							}
							
						});
					}
					else
						this.playerIsCzar.set(false);
				}
				// When joining table, client should send a player delta with a 0 Id and message "join".
				// Server should send a reply delta with next Id and message "you"
				// followed by zero or more player deltas that are the other people
			} else if (c == ActionDelta.class) {
				//TODO: Implement this type of delta.
			}
		}
	}

	public void showWhiteCardChooser(final Collection<Pair<Integer, String>> cards, final Context context) {
		cahActivity.runOnUiThread(new Runnable() {

			@Override
			public void run() {
				if(czarDialog != null) 
					czarDialog.dismiss();
				HorizontalScrollView scroller = new HorizontalScrollView(context);
				LinearLayout cardContainer = new LinearLayout(context);
				for(final Pair<Integer, String> card : cards) {
					CardView cv = new CardView(context);
					cv.setCardString(card.second);
					cv.setTextColor(Color.BLACK);
					cv.setCardColor(Color.WHITE);

					LayoutParams lp = new LayoutParams((int) (235* (context.getResources().getDisplayMetrics().densityDpi/160.)), (int) (300* (context.getResources().getDisplayMetrics().densityDpi/160.)));
					cv.setLayoutParams(lp);
					cv.setOnClickListener(new OnClickListener() {
						@Override
						public void onClick(View arg0) {
							czarChoseCard(card.first, card.second);
						}
					});
					cardContainer.addView(cv);
				}
				AlertDialog.Builder builder = new AlertDialog.Builder(context);
				builder.setTitle("Choose wining card");
				scroller.addView(cardContainer);
				builder.setView(scroller);
				czarDialog = builder.show();
			}
			
		});
		
	}

	protected void czarChoseCard(int winnerID, String cardText) {
		//TODO: Send message to server saying which card won.
		try {
			client.outgoing.put(new DeckDelta(winnerID, "winner", "hand", new Card[] {new Card(Card.Color.WHITE, cardText )}));
		} catch (InterruptedException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
		czarDialog.dismiss();
	}

	// Helper thread to handle incoming messages in a blocking fashion
	private class HandleMessageThread implements Runnable {
		public void run() {
			try {
				handleIncomingMessages(client.incoming, client.outgoing);
			} catch (Exception e) {
				System.err
						.println("Unexpected Incoming Message Handler Thread Exception: "
								+ e + "\n" + Arrays.toString(e.getStackTrace()));
			} finally {
				// Shutdown threads in case this thread encountered an error
				shutdown();
			}
		}
	}

	public class Player {
		public final int id;
		public final int imageID;
		public boolean czar;


		public Player(int playerId, int imageID, boolean isCzar) {
			this.imageID = imageID;
			this.id = playerId;
			this.czar = isCzar;
		}
	}

	public void shutdown() {
		go.set(false);
		messageHandler.interrupt();
		client.shutdown();
	}

	/**
	 * This function is called by CahClient when the server says that a card
	 * should be added to our hand in the UI.
	 *
	 * @param card
	 *            Card to be added to hand
	 */
	public void addCardToHand(final Card card) {
		cahActivity.runOnUiThread(new Runnable() {
			@Override
			public void run() {
				cahActivity.addCardToHand(card);
			}
		});
	}

	public void playCard(final Card card) {
		// lock deck
		cahActivity.playerCanPlayCard(false);
		
		Cah.performOnBackgroundThread(new Runnable() {
			@Override
			public void run() {
				try {
					client.outgoing.put(new DeckDelta(playerId, "play", "hand", new Card[] { card }));
				} catch (InterruptedException e) {
					// TODO Auto-generated catch block
					e.printStackTrace();
				}
			}
		});
	}

	/**
	 * This function is called by CahClient when a player joins our table. This
	 * adds the player to the table UI.
	 * 
	 * Added differing player icons
	 *
	 * @param player
	 *            The player that joined
	 */
	public void playerJoined(int id, final boolean isCzar) {
		// TODO: Implement this function.
		currentList.add(new Player(id, imageID, isCzar));
		cahActivity.runOnUiThread(new Runnable() {

			@Override
			public void run() {

				GameTable table = (GameTable) cahActivity.findViewById(R.id.gameTable);
				Bitmap playerBitmap;

				//gives each player a unique user image
				switch (imageID) {
					case 0:  playerBitmap = BitmapFactory.decodeResource(cahActivity.getResources(), R.drawable.cartman_tiny);
							 imageID++;
	                		 break;
		            case 1:  playerBitmap = BitmapFactory.decodeResource(cahActivity.getResources(), R.drawable.hulk);
		                     imageID++;
		            		 break;
		            case 2:  playerBitmap = BitmapFactory.decodeResource(cahActivity.getResources(), R.drawable.superman_color);
                    		 imageID++;
                    		 break;
		            case 3:  playerBitmap = BitmapFactory.decodeResource(cahActivity.getResources(), R.drawable.kenny_tiny);
			           		 imageID++;
			           		 break;
		            case 4:  playerBitmap = BitmapFactory.decodeResource(cahActivity.getResources(), R.drawable.butters_tiny);
			           		 imageID++;
			           		 break;
		            default:  playerBitmap = BitmapFactory.decodeResource(cahActivity.getResources(), R.drawable.droid_small);
			           		 imageID=0;
			           		 break;
				}
				table.addPlayerToTable(playerBitmap, isCzar);
		//        ImageView image = (ImageView) findViewById(R.id.test_image);
		//        image.setImageResource(R.drawable.test2);
			}
		});
	}

	/**
	 * This function is called by CahClient when a player leaves our table. This
	 * removes the player from the table UI.
	 *
	 * @param player
	 *            The player that left
	 */
	public void playerLeft(int id) {
		// TODO: Implement this function.
		//loop through currentList to find correct id and delete that member from the list
		for (int i = 0; i < currentList.size(); i++) {
			Player player = currentList.get(i);
			if (player.id == id) {
				currentList.remove(i);
				break;
			}
			else { System.out.println("Player does not exist in list");}
		}
	}

	/**
	 * This function is called by CahClient when the current czar changes. This
	 * changes the table UI so that a crown appears next to whoever is the czar.
	 *
	 * @param player
	 *            The player that is now the czar
	 */
	public void playerBecomesCzar(int id) {
		// TODO: Implement this function.
		//loop through list to find the correct id and change the isCzar variable for that Player object to be true
		for (int i = 0; i < currentList.size(); i++) {
			Player player = currentList.get(i);
			if (player.id == id) {
				player.czar = true;
			}
			else { System.out.println("Player does not exist in list");}
		}
		czarWhiteCards.clear();
	}

	public void showError(final String error, final String errorMessage) {
		cahActivity.runOnUiThread(new Runnable() {

			@Override
			public void run() {
				AlertDialog.Builder alertBuilder = new AlertDialog.Builder(
						cahActivity);
				alertBuilder.setTitle(error);
				alertBuilder.setMessage(errorMessage);
				alertBuilder.setPositiveButton("Close", null);
				alertBuilder.show();
			}
		});
	}

	public void showDebugText(final String debugText) {
		cahActivity.runOnUiThread(new Runnable() {
			@Override
			public void run() {
				TextView debugTextView = (TextView) cahActivity
						.findViewById(R.id.debugTextView);
				debugTextView.setText("Table ID: " + CahPlayer.this.tableID
						+ "\n" + debugText);
			}
		});
	}
}
