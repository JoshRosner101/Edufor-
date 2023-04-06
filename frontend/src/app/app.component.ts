import { Component, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Emitters } from './emitters/emitters';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  title = 'angular';

  constructor(
    private httpClient: HttpClient
  ) {}

  ngOnInit(): void {
    this.httpClient.get('/backend/users/user', {withCredentials: true}).subscribe(
      (res: any) => {
        Emitters.authEmitter.emit(true);
      },
      err => {
        Emitters.authEmitter.emit(false);
      }
    );
  }
}
