using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Graphics;
using mopacman.Components;
using mopacman.Controllers;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace mopacman.Animations
{
    class DirectionalAnimation<T>
        where T : Sprite
    {
        public event EventHandler Finished;

        private float   velocity;
        private float   duration;
        private float   remaining;
        private Vector2 direction;
        private T       component;
        private bool    finished;
        private Vector2 destiny;
                
        public DirectionalAnimation(T c, float duration )
        {
            this.component = c;
            this.duration = duration;
            this.velocity = 1f;
            this.finished = true;
        }

        public void Start(EnumDirections d, float distance)
        {
            this.destiny = destiny;
            this.finished = false;
            this.remaining = this.duration;

            switch (d)
            {
                case EnumDirections.North:
                    this.direction = new Vector2(0.0f, (distance * this.velocity) * -1.0f);
                    break;
                    
                case EnumDirections.South:
                    this.direction = new Vector2(0.0f, (distance * this.velocity));
                    break;

                case EnumDirections.East:
                    this.direction = new Vector2((distance * this.velocity), 0.0f);
                    break;

                case EnumDirections.West:
                    this.direction = new Vector2((distance * this.velocity) * -1.0f, 0.0f);
                    break;

                default:
                    break;
            }
        }

        public void Update(GameTime gameTime)
        {
            //if (!finished)
            //{
                if (this.remaining > 0.0f)
                {
                    var elapsed = (float)gameTime.ElapsedGameTime.TotalSeconds;
                    var position = this.component.Position;

                    float x = (position.X + (this.direction.X * elapsed));
                    float y = (position.Y + (this.direction.Y * elapsed));

                    this.component.SetPosition(x,y);

                    this.remaining -= elapsed;
                }
                else
                {
                    this.remaining = 0.0f;
                    this.finished = true;

                    if (this.Finished != null)
                        this.Finished.Invoke(this, null);
                }
            //}
        }
    }
}
