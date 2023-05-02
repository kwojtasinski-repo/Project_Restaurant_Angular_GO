import { Component, OnInit } from '@angular/core';
import { NavigationStart, Router } from '@angular/router';
import { take } from "rxjs";
import { AuthService } from './services/auth.service';
import { HeaderService } from './services/header.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
}
