package com.cah;

import android.app.Activity;
import android.content.Intent;
import android.os.Bundle;
import android.view.Menu;
import android.view.View;
import android.view.View.OnClickListener;
import android.widget.Button;
import android.widget.EditText;

public class GameLauncher extends Activity {

	@Override
	protected void onCreate(Bundle savedInstanceState) {
		super.onCreate(savedInstanceState);
		setContentView(R.layout.activity_game_launcher);
		
		final Button createNewTableButton = (Button) findViewById(R.id.buttonCreateNewTable);
		
		final Button joinExistingTableButton = (Button) findViewById(R.id.buttonJoinExistingTable);
		final EditText tableIdEditText = (EditText) findViewById(R.id.editTextTableId);
		
		final Button showCzarCardButton = (Button) findViewById(R.id.buttonShowCzarCard);
		final EditText czarCardText = (EditText) findViewById(R.id.czarCardText);
		
		createNewTableButton.setOnClickListener(new OnClickListener() {
			@Override
			public void onClick(View button) {
				Intent intent = new Intent(button.getContext(), Cah.class);
				intent.putExtra("COMMAND", "new");
				startActivity(intent);
			}
		});
		
		joinExistingTableButton.setOnClickListener(new OnClickListener() {
			@Override
			public void onClick(View v) {
				Intent intent = new Intent(v.getContext(), Cah.class);
				intent.putExtra("COMMAND", "join");
				intent.putExtra("TABLE_ID", tableIdEditText.getText().toString());
				startActivity(intent);
			}
		});
		
		showCzarCardButton.setOnClickListener(new OnClickListener() {
			@Override
			public void onClick(View v) {
				Intent intent = new Intent(v.getContext(), CzarActivity.class);
				intent.putExtra("CARD_TEXT", czarCardText.getText());
				startActivity(intent);
			}
		});
	}

	@Override
	public boolean onCreateOptionsMenu(Menu menu) {
		// Inflate the menu; this adds items to the action bar if it is present.
		getMenuInflater().inflate(R.menu.activity_game_launcher, menu);
		return true;
	}

}
