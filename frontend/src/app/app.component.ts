import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs'

interface IThreadItem {
  id: string,
  username: string,
  title: string,
  body: string
}

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css'],
})
export class AppComponent implements OnInit {
  public id = ''
  public username = ''
  public title = ''
  public body = ''

  public threadItems: IThreadItem[] = []

  constructor(
    private httpClient: HttpClient
  ) {}

  async ngOnInit() {
    await this.loadNewItems()
  }

  async loadNewItems() {
    this.threadItems = await lastValueFrom(this.httpClient.get<IThreadItem[]>('/backend/threads'))
  }

  async addThread() {
    await lastValueFrom(this.httpClient.post('/backend/threads', {
      id: this.id,
      username: this.username,
      title: this.title,
      body: this.body
    }))

    await this.loadNewItems()
    
    this.id = ''
    this.username = ''
    this.title = ''
    this.body = ''
    
  }
}
