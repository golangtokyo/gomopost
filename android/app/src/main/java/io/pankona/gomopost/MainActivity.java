package io.pankona.gomopost;

import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.view.View;
import android.webkit.WebView;
import android.webkit.WebViewClient;
import android.widget.Button;
import android.widget.EditText;

import gomopost.*;

public class MainActivity extends AppCompatActivity implements View.OnClickListener {

    private Button mGo;
    private Button mSend;
    private EditText mAddress;
    private EditText mName;
    private EditText mMessage;
    private WebView mLogView;
    private GomoPoster poster;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        mGo = findViewById(R.id.go);
        mGo.setOnClickListener(this);

        mSend = findViewById(R.id.send);
        mSend.setOnClickListener(this);

        mAddress = findViewById(R.id.address);
        mName = findViewById(R.id.name);
        mMessage = findViewById(R.id.message);

        mLogView = findViewById(R.id.webview);
        mLogView.getSettings().setJavaScriptEnabled(true);
        mLogView.setWebViewClient(new WebViewClient());
    }

    @Override
    public void onClick(View view) {
        switch (view.getId()) {
            case R.id.go:
                mLogView.post(new Runnable() {
                    @Override
                    public void run() {
                        String url = mAddress.getText().toString();
                        mLogView.loadUrl(url);
                    }
                });
                break;

            case R.id.send:
                poster = Gomopost.newClient(mAddress.getText().toString());
                String name = mName.getText().toString();
                String message = mMessage.getText().toString();
                try {
                    poster.post(name, message);
                } catch (Exception ex) {
                    ex.printStackTrace();
                }
                break;
        }
    }
}
