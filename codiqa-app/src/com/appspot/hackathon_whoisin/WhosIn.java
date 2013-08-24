package com.appspot.hackathon_whosin;

import android.app.Activity;
import android.os.Bundle;
import org.apache.cordova.*;

public class WhosIn extends DroidGap
{
    @Override
    public void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        super.loadUrl("https://hackathon-whosin.appspot.com/static/www/index.html");
    }
}

