import { Component, EventEmitter, OnInit } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Emitters } from '.././emitters/emitters';
import { Router } from '@angular/router';

@Component({
  selector: 'app-taskbar',
  templateUrl: './taskbar.component.html',
  styleUrls: ['./taskbar.component.css']
})
export class TaskbarComponent {

  authenticated = false;

  constructor(
    private httpClient: HttpClient,
    private router: Router
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
      this.router.routeReuseStrategy.shouldReuseRoute = () => false;
      this.router.onSameUrlNavigation = 'reload';
      this.router.navigate(['/home']);
    });
  }
}
