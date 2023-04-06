import { Component, EventEmitter, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Emitters } from '.././emitters/emitters';

@Component({
  selector: 'app-taskbar',
  templateUrl: './taskbar.component.html',
  styleUrls: ['./taskbar.component.css']
})
export class TaskbarComponent {

  authenticated = false;

  constructor(
    private httpClient: HttpClient
  ) {}
  
  async ngOnInit() {
    await Emitters.authEmitter.subscribe(
      (auth:boolean) => {
        this.authenticated = auth;
      }
    );
  }

  logout(): void {
    this.httpClient.post('/backend/users/logout', {}, {withCredentials: true}).subscribe(() => {
      this.authenticated = false;
    });
  }
}
