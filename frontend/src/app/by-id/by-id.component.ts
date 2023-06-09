import { OnInit, Component } from '@angular/core';
import { HttpClient } from '@angular/common/http'
import { lastValueFrom } from 'rxjs'
import { ActivatedRoute } from '@angular/router'
import { Title } from "@angular/platform-browser";
import { Router } from '@angular/router';

interface IReplyItem {
  replyid: number,
  username: string,
  body: string,
  time: string,
  replypost: number
}

//the variable names are the same ones in the databse
interface IThreadItem {
  id: number,
  username: string,
  title: string,
  body: string,
  time: string,
  replies: IReplyItem[]
}

@Component({
  selector: 'app-by-id',
  templateUrl: './by-id.component.html',
  styleUrls: ['./by-id.component.css']
})

export class ByIDComponent implements OnInit{
  authenticated = false;
  editPost = false;
  public username = ''
  public title = ''
  public body = ''
  public editedText = ''

  public threadItem: IThreadItem = {id: 0, username: "", title: "", body: "", time: "", replies: []}

  constructor(
    private httpClient: HttpClient,
    private route: ActivatedRoute,
    private titleService: Title,
    private router: Router
  ) {
    this.titleService.setTitle("Thread #"+ this.route.snapshot.paramMap.get('id'));
  }

  async ngOnInit() {
    await this.loadNewItems()
    this.httpClient.get('/backend/users/user', {withCredentials: true}).subscribe(
      (res: any) => {
        this.authenticated = true;
        this.username = res.name;
      },
      err => {
        this.authenticated = false;
      }
    );
  }

  async loadNewItems() {
    const urlpath = "/backend/threads/" + this.route.snapshot.paramMap.get('id')
    this.threadItem = await lastValueFrom(this.httpClient.get<IThreadItem>(urlpath))
  }

  async addReply() {
    let currentID = this.route.snapshot.paramMap.get('id')
    const urlpath = "/backend/threads/" + currentID
    await lastValueFrom(this.httpClient.post(urlpath, {
      username: this.username,
      body: this.body,
      time: new Date().toLocaleString("en-US"),
      replypost: Number(currentID)
    }))

    await this.loadNewItems()
    console.log(this.username)
    
    this.username = ''
    this.body = ''
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['/home/'+ currentID])
  }

  async delete() {
    let currentID = this.route.snapshot.paramMap.get('id')
    const urlpath = "/backend/threads/" + currentID
    await lastValueFrom(this.httpClient.delete(urlpath))
    this.router.navigate(['/home'])
  }

  allowEdit() {
    this.editPost = true;
  }

  async edit() {
    let currentID = this.route.snapshot.paramMap.get('id')
    const urlpath = "/backend/threads/" + currentID
    
    await lastValueFrom(this.httpClient.put(urlpath, {
      id: Number(currentID),
      username: this.threadItem.username,
      title: this.threadItem.title,
      body: this.editedText,
      time: this.threadItem.time
    }))
    this.router.routeReuseStrategy.shouldReuseRoute = () => false;
    this.router.onSameUrlNavigation = 'reload';
    this.router.navigate(['/home/'+ currentID])
  }
}
