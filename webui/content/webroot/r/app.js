(()=>{"use strict";var t=function(){},s=function(){function t(t){this.state=t;var s=window.location,e=("https:"===s.protocol?"wss:":"ws:")+"//"+s.host+"/ws/";this.ws=new WebSocket(e),this.ws.onmessage=this.onmessage}return t.prototype.onmessage=function(t){var s=JSON.parse(t.data);switch(s.c){case"vmu":this.state.viewModel=s.d}},t}();document.addEventListener("load",(function(e){var n=new t;new s(n)}))})();
//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly93ZWJ1aS8uL3NyYy9pbmRleC50cyJdLCJuYW1lcyI6WyJzdGF0ZSIsInRoaXMiLCJ3aW5kb3ciLCJsb2NhdGlvbiIsInVybCIsIndzIiwiV2ViU29ja2V0Iiwib25tZXNzYWdlIiwiZSIsIm1zZyIsIkpTT04iLCJwYXJzZSIsImRhdGEiLCJjIiwidmlld01vZGVsIiwiZCIsImRvY3VtZW50IiwiYWRkRXZlbnRMaXN0ZW5lciIsImV2IiwiU3RhdGUiLCJIb3N0Il0sIm1hcHBpbmdzIjoibUJBR0EsbUJBSUEsYUFJSSxXQUFZQSxHQUNSQyxLQUFLRCxNQUFRQSxFQUVQLE1BQW1CRSxPQUFPQyxTQUMxQkMsR0FBb0IsV0FEWCxXQUNzQixPQUFTLE9BQVMsS0FEbEMsT0FDZ0QsT0FFckVILEtBQUtJLEdBQUssSUFBSUMsVUFBVUYsR0FDeEJILEtBQUtJLEdBQUdFLFVBQVlOLEtBQUtNLFVBV2pDLE9BUkksWUFBQUEsVUFBQSxTQUFVQyxHQUNOLElBQUlDLEVBQU1DLEtBQUtDLE1BQU1ILEVBQUVJLE1BQ3ZCLE9BQVFILEVBQUlJLEdBQ1IsSUFBSyxNQUNEWixLQUFLRCxNQUFNYyxVQUFZTCxFQUFJTSxJQUkzQyxFQXRCQSxHQXdCQUMsU0FBU0MsaUJBQWlCLFFBQVEsU0FBQUMsR0FDOUIsSUFBSWxCLEVBQVEsSUFBSW1CLEVBQ0wsSUFBSUMsRUFBS3BCLE8iLCJmaWxlIjoiYXBwLmpzIiwic291cmNlc0NvbnRlbnQiOlsiaW1wb3J0IHtWaWV3TW9kZWx9IGZyb20gJy4vdmlld21vZGVsJztcbmltcG9ydCB7TzJJbmNvbWluZ01lc3NhZ2V9IGZyb20gJy4vbWVzc2FnZXMnO1xuXG5jbGFzcyBTdGF0ZSB7XG4gICAgcHVibGljIHZpZXdNb2RlbDogVmlld01vZGVsO1xufVxuXG5jbGFzcyBIb3N0IHtcbiAgICBwcml2YXRlIHN0YXRlOiBTdGF0ZTtcbiAgICBwcml2YXRlIHdzOiBXZWJTb2NrZXQ7XG5cbiAgICBjb25zdHJ1Y3RvcihzdGF0ZTogU3RhdGUpIHtcbiAgICAgICAgdGhpcy5zdGF0ZSA9IHN0YXRlO1xuXG4gICAgICAgIGNvbnN0IHtwcm90b2NvbCwgaG9zdH0gPSB3aW5kb3cubG9jYXRpb247XG4gICAgICAgIGNvbnN0IHVybCA9IChwcm90b2NvbCA9PT0gXCJodHRwczpcIiA/IFwid3NzOlwiIDogXCJ3czpcIikgKyBcIi8vXCIgKyBob3N0ICsgXCIvd3MvXCI7XG5cbiAgICAgICAgdGhpcy53cyA9IG5ldyBXZWJTb2NrZXQodXJsKTtcbiAgICAgICAgdGhpcy53cy5vbm1lc3NhZ2UgPSB0aGlzLm9ubWVzc2FnZTtcbiAgICB9XG5cbiAgICBvbm1lc3NhZ2UoZTogTWVzc2FnZUV2ZW50PHN0cmluZz4pIHtcbiAgICAgICAgbGV0IG1zZyA9IEpTT04ucGFyc2UoZS5kYXRhKSBhcyBPMkluY29taW5nTWVzc2FnZTtcbiAgICAgICAgc3dpdGNoIChtc2cuYykge1xuICAgICAgICAgICAgY2FzZSBcInZtdVwiOiAvLyB2aWV3LW1vZGVsIHVwZGF0ZVxuICAgICAgICAgICAgICAgIHRoaXMuc3RhdGUudmlld01vZGVsID0gbXNnLmQ7XG4gICAgICAgICAgICAgICAgYnJlYWs7XG4gICAgICAgIH1cbiAgICB9XG59XG5cbmRvY3VtZW50LmFkZEV2ZW50TGlzdGVuZXIoXCJsb2FkXCIsIGV2ID0+IHtcbiAgICBsZXQgc3RhdGUgPSBuZXcgU3RhdGUoKTtcbiAgICBsZXQgaG9zdCA9IG5ldyBIb3N0KHN0YXRlKTtcbn0pO1xuIl0sInNvdXJjZVJvb3QiOiIifQ==