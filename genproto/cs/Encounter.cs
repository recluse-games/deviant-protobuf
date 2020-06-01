// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: Encounter.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace Deviant {

  /// <summary>Holder for reflection information generated from Encounter.proto</summary>
  public static partial class EncounterReflection {

    #region Descriptor
    /// <summary>File descriptor for Encounter.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static EncounterReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "Cg9FbmNvdW50ZXIucHJvdG8SB0RldmlhbnQaD0FsaWdubWVudC5wcm90bxoL",
            "Qm9hcmQucHJvdG8aDEVudGl0eS5wcm90bxoKVHVybi5wcm90byLWAQoJRW5j",
            "b3VudGVyEgoKAmlkGAEgASgJEhEKCWNvbXBsZXRlZBgCIAEoCBIdCgVib2Fy",
            "ZBgDIAEoCzIOLkRldmlhbnQuQm9hcmQSGwoEdHVybhgEIAEoCzINLkRldmlh",
            "bnQuVHVybhIZChFhY3RpdmVFbnRpdHlPcmRlchgFIAMoCRIlCgxhY3RpdmVF",
            "bnRpdHkYBiABKAsyDy5EZXZpYW50LkVudGl0eRIsChB3aW5uaW5nQWxpZ25t",
            "ZW50GAcgASgOMhIuRGV2aWFudC5BbGlnbm1lbnRCCVoHZGV2aWFudGIGcHJv",
            "dG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::Deviant.AlignmentReflection.Descriptor, global::Deviant.BoardReflection.Descriptor, global::Deviant.EntityReflection.Descriptor, global::Deviant.TurnReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, new pbr::GeneratedClrTypeInfo[] {
            new pbr::GeneratedClrTypeInfo(typeof(global::Deviant.Encounter), global::Deviant.Encounter.Parser, new[]{ "Id", "Completed", "Board", "Turn", "ActiveEntityOrder", "ActiveEntity", "WinningAlignment" }, null, null, null, null)
          }));
    }
    #endregion

  }
  #region Messages
  public sealed partial class Encounter : pb::IMessage<Encounter> {
    private static readonly pb::MessageParser<Encounter> _parser = new pb::MessageParser<Encounter>(() => new Encounter());
    private pb::UnknownFieldSet _unknownFields;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pb::MessageParser<Encounter> Parser { get { return _parser; } }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public static pbr::MessageDescriptor Descriptor {
      get { return global::Deviant.EncounterReflection.Descriptor.MessageTypes[0]; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    pbr::MessageDescriptor pb::IMessage.Descriptor {
      get { return Descriptor; }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Encounter() {
      OnConstruction();
    }

    partial void OnConstruction();

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Encounter(Encounter other) : this() {
      id_ = other.id_;
      completed_ = other.completed_;
      board_ = other.board_ != null ? other.board_.Clone() : null;
      turn_ = other.turn_ != null ? other.turn_.Clone() : null;
      activeEntityOrder_ = other.activeEntityOrder_.Clone();
      activeEntity_ = other.activeEntity_ != null ? other.activeEntity_.Clone() : null;
      winningAlignment_ = other.winningAlignment_;
      _unknownFields = pb::UnknownFieldSet.Clone(other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public Encounter Clone() {
      return new Encounter(this);
    }

    /// <summary>Field number for the "id" field.</summary>
    public const int IdFieldNumber = 1;
    private string id_ = "";
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public string Id {
      get { return id_; }
      set {
        id_ = pb::ProtoPreconditions.CheckNotNull(value, "value");
      }
    }

    /// <summary>Field number for the "completed" field.</summary>
    public const int CompletedFieldNumber = 2;
    private bool completed_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Completed {
      get { return completed_; }
      set {
        completed_ = value;
      }
    }

    /// <summary>Field number for the "board" field.</summary>
    public const int BoardFieldNumber = 3;
    private global::Deviant.Board board_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Board Board {
      get { return board_; }
      set {
        board_ = value;
      }
    }

    /// <summary>Field number for the "turn" field.</summary>
    public const int TurnFieldNumber = 4;
    private global::Deviant.Turn turn_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Turn Turn {
      get { return turn_; }
      set {
        turn_ = value;
      }
    }

    /// <summary>Field number for the "activeEntityOrder" field.</summary>
    public const int ActiveEntityOrderFieldNumber = 5;
    private static readonly pb::FieldCodec<string> _repeated_activeEntityOrder_codec
        = pb::FieldCodec.ForString(42);
    private readonly pbc::RepeatedField<string> activeEntityOrder_ = new pbc::RepeatedField<string>();
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public pbc::RepeatedField<string> ActiveEntityOrder {
      get { return activeEntityOrder_; }
    }

    /// <summary>Field number for the "activeEntity" field.</summary>
    public const int ActiveEntityFieldNumber = 6;
    private global::Deviant.Entity activeEntity_;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Entity ActiveEntity {
      get { return activeEntity_; }
      set {
        activeEntity_ = value;
      }
    }

    /// <summary>Field number for the "winningAlignment" field.</summary>
    public const int WinningAlignmentFieldNumber = 7;
    private global::Deviant.Alignment winningAlignment_ = global::Deviant.Alignment.Friendly;
    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public global::Deviant.Alignment WinningAlignment {
      get { return winningAlignment_; }
      set {
        winningAlignment_ = value;
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override bool Equals(object other) {
      return Equals(other as Encounter);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public bool Equals(Encounter other) {
      if (ReferenceEquals(other, null)) {
        return false;
      }
      if (ReferenceEquals(other, this)) {
        return true;
      }
      if (Id != other.Id) return false;
      if (Completed != other.Completed) return false;
      if (!object.Equals(Board, other.Board)) return false;
      if (!object.Equals(Turn, other.Turn)) return false;
      if(!activeEntityOrder_.Equals(other.activeEntityOrder_)) return false;
      if (!object.Equals(ActiveEntity, other.ActiveEntity)) return false;
      if (WinningAlignment != other.WinningAlignment) return false;
      return Equals(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override int GetHashCode() {
      int hash = 1;
      if (Id.Length != 0) hash ^= Id.GetHashCode();
      if (Completed != false) hash ^= Completed.GetHashCode();
      if (board_ != null) hash ^= Board.GetHashCode();
      if (turn_ != null) hash ^= Turn.GetHashCode();
      hash ^= activeEntityOrder_.GetHashCode();
      if (activeEntity_ != null) hash ^= ActiveEntity.GetHashCode();
      if (WinningAlignment != global::Deviant.Alignment.Friendly) hash ^= WinningAlignment.GetHashCode();
      if (_unknownFields != null) {
        hash ^= _unknownFields.GetHashCode();
      }
      return hash;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public override string ToString() {
      return pb::JsonFormatter.ToDiagnosticString(this);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void WriteTo(pb::CodedOutputStream output) {
      if (Id.Length != 0) {
        output.WriteRawTag(10);
        output.WriteString(Id);
      }
      if (Completed != false) {
        output.WriteRawTag(16);
        output.WriteBool(Completed);
      }
      if (board_ != null) {
        output.WriteRawTag(26);
        output.WriteMessage(Board);
      }
      if (turn_ != null) {
        output.WriteRawTag(34);
        output.WriteMessage(Turn);
      }
      activeEntityOrder_.WriteTo(output, _repeated_activeEntityOrder_codec);
      if (activeEntity_ != null) {
        output.WriteRawTag(50);
        output.WriteMessage(ActiveEntity);
      }
      if (WinningAlignment != global::Deviant.Alignment.Friendly) {
        output.WriteRawTag(56);
        output.WriteEnum((int) WinningAlignment);
      }
      if (_unknownFields != null) {
        _unknownFields.WriteTo(output);
      }
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public int CalculateSize() {
      int size = 0;
      if (Id.Length != 0) {
        size += 1 + pb::CodedOutputStream.ComputeStringSize(Id);
      }
      if (Completed != false) {
        size += 1 + 1;
      }
      if (board_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Board);
      }
      if (turn_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(Turn);
      }
      size += activeEntityOrder_.CalculateSize(_repeated_activeEntityOrder_codec);
      if (activeEntity_ != null) {
        size += 1 + pb::CodedOutputStream.ComputeMessageSize(ActiveEntity);
      }
      if (WinningAlignment != global::Deviant.Alignment.Friendly) {
        size += 1 + pb::CodedOutputStream.ComputeEnumSize((int) WinningAlignment);
      }
      if (_unknownFields != null) {
        size += _unknownFields.CalculateSize();
      }
      return size;
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(Encounter other) {
      if (other == null) {
        return;
      }
      if (other.Id.Length != 0) {
        Id = other.Id;
      }
      if (other.Completed != false) {
        Completed = other.Completed;
      }
      if (other.board_ != null) {
        if (board_ == null) {
          Board = new global::Deviant.Board();
        }
        Board.MergeFrom(other.Board);
      }
      if (other.turn_ != null) {
        if (turn_ == null) {
          Turn = new global::Deviant.Turn();
        }
        Turn.MergeFrom(other.Turn);
      }
      activeEntityOrder_.Add(other.activeEntityOrder_);
      if (other.activeEntity_ != null) {
        if (activeEntity_ == null) {
          ActiveEntity = new global::Deviant.Entity();
        }
        ActiveEntity.MergeFrom(other.ActiveEntity);
      }
      if (other.WinningAlignment != global::Deviant.Alignment.Friendly) {
        WinningAlignment = other.WinningAlignment;
      }
      _unknownFields = pb::UnknownFieldSet.MergeFrom(_unknownFields, other._unknownFields);
    }

    [global::System.Diagnostics.DebuggerNonUserCodeAttribute]
    public void MergeFrom(pb::CodedInputStream input) {
      uint tag;
      while ((tag = input.ReadTag()) != 0) {
        switch(tag) {
          default:
            _unknownFields = pb::UnknownFieldSet.MergeFieldFrom(_unknownFields, input);
            break;
          case 10: {
            Id = input.ReadString();
            break;
          }
          case 16: {
            Completed = input.ReadBool();
            break;
          }
          case 26: {
            if (board_ == null) {
              Board = new global::Deviant.Board();
            }
            input.ReadMessage(Board);
            break;
          }
          case 34: {
            if (turn_ == null) {
              Turn = new global::Deviant.Turn();
            }
            input.ReadMessage(Turn);
            break;
          }
          case 42: {
            activeEntityOrder_.AddEntriesFrom(input, _repeated_activeEntityOrder_codec);
            break;
          }
          case 50: {
            if (activeEntity_ == null) {
              ActiveEntity = new global::Deviant.Entity();
            }
            input.ReadMessage(ActiveEntity);
            break;
          }
          case 56: {
            WinningAlignment = (global::Deviant.Alignment) input.ReadEnum();
            break;
          }
        }
      }
    }

  }

  #endregion

}

#endregion Designer generated code
