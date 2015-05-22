using Microsoft.Xna.Framework;
using Microsoft.Xna.Framework.Audio;
using Microsoft.Xna.Framework.Graphics;
using Microsoft.Xna.Framework.Input;
using mopacman.Components;
using mopacman.Controllers;
using mopacman.Scenes;
using System;
using System.Collections.Generic;

namespace mopacman
{
    /// <summary>
    /// This is the main type for your game.
    /// </summary>
    public class MyGame : Game
    {
        public enum MyGameStates
        {
            Welcome = 0,
            GamePlay = 1
        }

        public static GraphicsDeviceManager Graphics;
        public static Camera Camera;
        public static SpriteBatch SpriteBatch;
        
        Dictionary<MyGameStates,GameScene> scenes;
        MyGameStates currentGameState;
        SoundEffect soundTest;
                        
        public MyGame()
        {
            Graphics = new GraphicsDeviceManager(this);
            Content.RootDirectory = "Content";
        }

        /// <summary>
        /// Allows the game to perform any initialization it needs to before starting to run.
        /// This is where it can query for any required services and load any non-graphic
        /// related content.  Calling base.Initialize will enumerate through any components
        /// and initialize them as well.
        /// </summary>
        protected override void Initialize()
        {
            this.scenes = new Dictionary<MyGameStates, GameScene>();

            // TODO: Add your initialization logic here
            TitleScreenScene scene1 = new TitleScreenScene(this);
            MazeScene scene2 = new MazeScene(this);
                                                
            this.Components.Add(scene1);
            this.Components.Add(scene2);

            scene1.Visible = true;
            scene1.Enabled = true;
                        
            this.scenes.Add(MyGameStates.Welcome, scene1);
            this.scenes.Add(MyGameStates.GamePlay, scene2);

            this.currentGameState = MyGameStates.Welcome;

            base.Initialize();
        }

        /// <summary>
        /// LoadContent will be called once per game and is the place to load
        /// all of your content.
        /// </summary>
        protected override void LoadContent()
        {
            MyGame.Graphics.PreferredBackBufferWidth = 465;
            MyGame.Graphics.PreferredBackBufferHeight = 550;
            MyGame.Graphics.IsFullScreen = false;
            MyGame.Graphics.ApplyChanges();

            MyGame.Camera = new Camera(new Vector2(10, 50), MyGame.Graphics.PreferredBackBufferWidth);

            // Create a new SpriteBatch, which can be used to draw textures.
            MyGame.SpriteBatch = new SpriteBatch(GraphicsDevice);

            //Carregando arquivos de som independente de onde serão usados
            //this.soundTest = this.Content.Load<SoundEffect>("SoundEffects\\pacman_chomp");
        }

        /// <summary>
        /// UnloadContent will be called once per game and is the place to unload
        /// game-specific content.
        /// </summary>
        protected override void UnloadContent()
        {
            // TODO: Unload any non ContentManager content here
            this.Content.Unload();
        }

        /// <summary>
        /// Allows the game to run logic such as updating the world,
        /// checking for collisions, gathering input, and playing audio.
        /// </summary>
        /// <param name="gameTime">Provides a snapshot of timing values.</param>
        protected override void Update(GameTime gameTime)
        {
            KeyboardState state = Keyboard.GetState();

            //Saída
            if (GamePad.GetState(PlayerIndex.One).Buttons.Back == ButtonState.Pressed || state.IsKeyDown(Keys.Escape))
                Exit();

            if(state.IsKeyDown(Keys.A))
            {
                this.soundTest.Play();
            }

            //Fullscreen
            if (state.IsKeyDown(Keys.LeftAlt) && state.IsKeyDown(Keys.Enter))
            {
                MyGame.Graphics.IsFullScreen = !MyGame.Graphics.IsFullScreen;
                MyGame.Graphics.ApplyChanges();
            }
                        
            //Tela de apresentação para gameplay
            if (this.currentGameState == MyGameStates.Welcome)
            {
                if (state.IsKeyDown(Keys.Enter))
                {
                    this.scenes[MyGameStates.Welcome].End();
                    this.scenes[MyGameStates.GamePlay].Begin();

                    this.currentGameState = MyGameStates.GamePlay;
                }
            }
                
        
            // TODO: Add your update logic here
            base.Update(gameTime);
        }

        /// <summary>
        /// This is called when the game should draw itself.
        /// </summary>
        /// <param name="gameTime">Provides a snapshot of timing values.</param>
        protected override void Draw(GameTime gameTime)
        {
            GraphicsDevice.Clear(Color.Black);

            MyGame.SpriteBatch.Begin();
            base.Draw(gameTime);
            MyGame.SpriteBatch.End();
        }
    }
}
